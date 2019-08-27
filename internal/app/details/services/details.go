package services

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"github.com/sdgmf/go-project-sample/internal/pkg/models"
	"github.com/sdgmf/go-project-sample/internal/app/details/repositories"
)

type DetailsService interface {
	Get(ID uint64) (*models.Detail, error)
}

type DefaultDetailsService struct {
	logger     *zap.Logger
	Repository repositories.DetailsRepository
}

func NewDetailService(logger *zap.Logger, Repository repositories.DetailsRepository) DetailsService {
	return &DefaultDetailsService{
		logger:  logger.With(zap.String("type","DefaultDetailsService")),
		Repository: Repository,
	}
}

func (s *DefaultDetailsService) Get(ID uint64) (p *models.Detail, err error) {
	if p, err = s.Repository.Get(ID); err != nil {
		return nil, errors.Wrap(err, "detail service get detail error")
	}

	return
}
