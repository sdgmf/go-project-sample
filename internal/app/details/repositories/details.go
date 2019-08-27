package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"github.com/sdgmf/go-project-sample/internal/pkg/models"
)

type DetailsRepository interface {
	Get(ID uint64) (p *models.Detail, err error)
}

type MysqlDetailsRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewMysqlDetailsRepository(logger *zap.Logger, db *gorm.DB) DetailsRepository {
	return &MysqlDetailsRepository{
		logger: logger.With(zap.String("type", "DetailsRepository")),
		db:     db,
	}
}

func (s *MysqlDetailsRepository) Get(ID uint64) (p *models.Detail, err error) {
	p = new(models.Detail)
	if err = s.db.Model(p).Where("id = ?", ID).First(p).Error; err != nil {
		return nil, errors.Wrapf(err, "get product error[id=%d]", ID)
	}
	return
}
