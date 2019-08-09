// +build wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-project-sample/internal/pkg/config"
	"github.com/sdgmf/go-project-sample/internal/pkg/database"
	"github.com/sdgmf/go-project-sample/internal/pkg/log"
	"github.com/sdgmf/go-project-sample/internal/app/ratings/services"
	"github.com/sdgmf/go-project-sample/internal/app/ratings/repositorys"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	//repositorys.ProviderSet,
	ProviderSet,
)


func CreateRatingsController(cf string, sto repositorys.RatingsRepository) (*RatingsController, error) {
	panic(wire.Build(testProviderSet))
}
