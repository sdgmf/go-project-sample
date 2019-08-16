package grpc

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/sdgmf/go-project-sample/internal/pkg/utils/netutil"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ServerOptions struct {
	Port int
}

func NewServerOptions(v *viper.Viper) (*ServerOptions, error) {
	var (
		err error
		o   = new(ServerOptions)
	)
	if err = v.UnmarshalKey("grpc", o); err != nil {
		return nil, err
	}

	return o, nil
}

type Server struct {
	o         *ServerOptions
	app       string
	host      string
	port      int
	logger    *zap.Logger
	server    *grpc.Server
	consulCli *consulApi.Client
}


type InitServers func(s *grpc.Server)

func NewServer(o *ServerOptions, logger *zap.Logger, init InitServers, consulCli *consulApi.Client, tracer opentracing.Tracer) (*Server, error) {
	// initialize grpc server
	var gs *grpc.Server
	logger = logger.With(zap.String("type", "grpc"))
	{
		grpc_prometheus.EnableHandlingTimeHistogram()
		gs = grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(),
				grpc_prometheus.StreamServerInterceptor,
				grpc_zap.StreamServerInterceptor(logger),
				grpc_recovery.StreamServerInterceptor(),
				otgrpc.OpenTracingStreamServerInterceptor(tracer),
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
				grpc_zap.UnaryServerInterceptor(logger),
				grpc_recovery.UnaryServerInterceptor(),
				otgrpc.OpenTracingServerInterceptor(tracer),
			)),
		)
		init(gs)
	}

	return &Server{
		o:         o,
		logger:    logger.With(zap.String("type", "grpc.Server")),
		server:    gs,
		consulCli: consulCli,
	}, nil
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

	s.logger.Info("grpc server starting ...", zap.String("addr", addr))
	go func() {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		if err := s.server.Serve(lis); err != nil {
			s.logger.Fatal("failed to serve: %v", zap.Error(err))
		}
	}()

	if err := s.register(); err != nil {
		return errors.Wrap(err, "register grpc server error")
	}

	return nil
}

func (s *Server) register() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	for key, _ := range s.server.GetServiceInfo() {
		check := &consulApi.AgentServiceCheck{
			Interval:                       "10s",
			DeregisterCriticalServiceAfter: "60m",
			TCP:                            addr,
		}

		id := fmt.Sprintf("%s[%s:%d]", key, s.host, s.port)

		svcReg := &consulApi.AgentServiceRegistration{
			ID:                id,
			Name:              key,
			Tags:              []string{"grpc"},
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
		s.logger.Info("register grpc service success", zap.String("id", id))
	}

	return nil
}

func (s *Server) deRegister() error {
	for key, _ := range s.server.GetServiceInfo() {
		id := fmt.Sprintf("%s[%s:%d]", key, s.host, s.port)

		err := s.consulCli.Agent().ServiceDeregister(id)
		if err != nil {
			return errors.Wrapf(err, "deregister service error[id=%s]", id)
		}
		s.logger.Info("deregister service success ", zap.String("id", id))
	}

	return nil
}

func (s *Server) Stop() error {
	s.logger.Info("grpc server stopping ...")
	if err := s.deRegister(); err != nil {
		return errors.Wrap(err, "deregister grpc server error")
	}
	s.server.GracefulStop()
	return nil
}
