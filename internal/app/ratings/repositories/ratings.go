package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sdgmf/go-project-sample/internal/pkg/models"
	"go.uber.org/zap"
)

type RatingsRepository interface {
	Get(productID uint64) (p *models.Rating, err error)
}

type MysqlRatingsRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewMysqlRatingsRepository(logger *zap.Logger, db *gorm.DB) RatingsRepository {
	return &MysqlRatingsRepository{
		logger: logger.With(zap.String("type", "RatingsRepository")),
		db:     db,
	}
}

func (s *MysqlRatingsRepository) Get(productID uint64) (p *models.Rating, err error) {
	p = new(models.Rating)
	if err = s.db.Model(p).Where("product_id = ?", productID).First(p).Error; err != nil {
		return nil, errors.Wrapf(err, "get rating error[productID=%d]", productID)
	}
	return
}
