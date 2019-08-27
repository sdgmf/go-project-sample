// +build wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-project-sample/internal/pkg/config"
	"github.com/sdgmf/go-project-sample/internal/pkg/database"
	"github.com/sdgmf/go-project-sample/internal/pkg/log"
	"github.com/sdgmf/go-project-sample/internal/app/details/services"
	"github.com/sdgmf/go-project-sample/internal/app/details/repositories"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	//repositories.ProviderSet,
	ProviderSet,
)


func CreateDetailsController(cf string, sto repositories.DetailsRepository) (*DetailsController, error) {
	panic(wire.Build(testProviderSet))
}
