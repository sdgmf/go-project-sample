// +build wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-project-sample/internal/pkg/config"
	"github.com/sdgmf/go-project-sample/internal/pkg/database"
	"github.com/sdgmf/go-project-sample/internal/pkg/log"
	"github.com/sdgmf/go-project-sample/internal/app/reviews/services"
	"github.com/sdgmf/go-project-sample/internal/app/reviews/repositorys"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	//repositorys.ProviderSet,
	ProviderSet,
)


func CreateReviewsController(cf string, sto repositorys.ReviewsRepository) (*ReviewsController, error) {
	panic(wire.Build(testProviderSet))
}
