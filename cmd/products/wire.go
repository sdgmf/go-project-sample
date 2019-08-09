// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-project-sample/internal/app/products"
	"github.com/sdgmf/go-project-sample/internal/app/products/controllers"
	"github.com/sdgmf/go-project-sample/internal/app/products/services"
	"github.com/sdgmf/go-project-sample/internal/app/products/grpcclients"
	"github.com/sdgmf/go-project-sample/internal/pkg/config"
	"github.com/sdgmf/go-project-sample/internal/pkg/consul"
	"github.com/sdgmf/go-project-sample/internal/pkg/log"
	"github.com/sdgmf/go-project-sample/internal/pkg/jaeger"
	"github.com/sdgmf/go-project-sample/internal/pkg/app"
	"github.com/sdgmf/go-project-sample/internal/pkg/transports/grpc"
	"github.com/sdgmf/go-project-sample/internal/pkg/transports/http"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	grpcclients.ProviderSet,
	controllers.ProviderSet,
	services.ProviderSet,
	products.ProviderSet,
)


func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
