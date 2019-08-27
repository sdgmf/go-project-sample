// +build wireinject

package repositories

import (
	"github.com/google/wire"
	"github.com/sdgmf/go-project-sample/internal/pkg/config"
	"github.com/sdgmf/go-project-sample/internal/pkg/database"
	"github.com/sdgmf/go-project-sample/internal/pkg/log"
)



var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateReviewRepository(f string) (ReviewsRepository, error) {
	panic(wire.Build(testProviderSet))
}

