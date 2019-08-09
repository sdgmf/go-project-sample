package http

import (
	"context"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sdgmf/go-project-sample/internal/pkg/transports/http/middlewares/ginprom"
	"github.com/sdgmf/go-project-sample/internal/pkg/utils/netutil"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Options struct {
	Port int
	Mode string
}

type Server struct {
	o          *Options
	app        string
	host       string
	port       int
	logger     *zap.Logger
	router     *gin.Engine
	httpServer http.Server
	consulCli  *consulApi.Client
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, err
	}

	return o, err
}

type InitControllers func(r *gin.Engine)

func NewRouter(o *Options, logger *zap.Logger, init InitControllers, tracer opentracing.Tracer) *gin.Engine {

	// 配置gin
	gin.SetMode(o.Mode)
	r := gin.New()

	r.Use(gin.Recovery()) // panic之后自动恢复
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(ginprom.New(r).Middleware()) // 添加prometheus 监控
	r.Use(ginhttp.Middleware(tracer))

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	pprof.Register(r)

	init(r)

	return r
}

func New(o *Options, logger *zap.Logger, router *gin.Engine, consulCli *consulApi.Client) (*Server, error) {
	var s = &Server{
		logger:    logger.With(zap.String("type", "http.Server")),
		router:    router,
		consulCli: consulCli,
		o:         o,
	}

	return s, nil
}

func (s *Server) Application(name string) {
	s.app = name
}

func (s *Server) Start() error {
	s.port = s.o.Port
	if s.port == 0 {
		s.port = netutil.GetAvailablePort()
	}

	s.host = netutil.GetLocalIP4()

	if s.host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	s.httpServer = http.Server{Addr: addr, Handler: s.router}

	s.logger.Info("http server starting ...", zap.String("addr", addr))
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("start http server err", zap.Error(err))
			return
		}
	}()

	if err := s.register(); err != nil {
		return errors.Wrap(err, "register http server error")
	}
	return nil
}

func (s *Server) register() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	check := &consulApi.AgentServiceCheck{
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "60m",
		TCP:                            addr,
	}

	id := fmt.Sprintf("%s[%s:%d]", s.app, s.host, s.port)

	svcReg := &consulApi.AgentServiceRegistration{
		ID:                id,
		Name:              string(s.app),
		Tags:              []string{"http"},
		Port:              s.port,
		Address:           s.host,
		EnableTagOverride: true,
		Check:             check,
		Checks:            nil,
	}

	err := s.consulCli.Agent().ServiceRegister(svcReg)
	if err != nil {
		return errors.Wrap(err, "register service error")
	}
	s.logger.Info("register http server success", zap.String("id", id))

	return nil
}

func (s *Server) deRegister() error {
	id := fmt.Sprintf("%s[%s:%d]", s.app, s.host, s.port)

	err := s.consulCli.Agent().ServiceDeregister(id)
	if err != nil {
		return errors.Wrapf(err, "deregister service error[key=%s]", id)
	}
	s.logger.Info("deregister service success ", zap.String("service", id))

	return nil
}

func (s *Server) Stop() error {
	s.logger.Info("stopping http server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) // 平滑关闭,等待5秒钟处理
	defer cancel()

	if err := s.deRegister(); err != nil {
		return errors.Wrap(err, "deregister http server error")
	}

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "shutdown http server error")
	}

	return nil
}

var ProviderSet = wire.NewSet(New, NewRouter, NewOptions)
