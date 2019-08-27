package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sdgmf/go-project-sample/internal/pkg/models"
	"go.uber.org/zap"
)

type ReviewsRepository interface {
	Query(productID uint64) (p []*models.Review, err error)
}

type MysqlReviewsRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewMysqlReviewsRepository(logger *zap.Logger, db *gorm.DB) ReviewsRepository {
	return &MysqlReviewsRepository{
		logger: logger.With(zap.String("type", "ReviewsRepository")),
		db:     db,
	}
}

func (s *MysqlReviewsRepository) Query(productID uint64) (rs []*models.Review, err error) {
	if err = s.db.Table("reviews").Where("product_id = ?", productID).Find(&rs).Error; err != nil {
		return nil, errors.Wrapf(err, "get review error[productID=%d]", productID)
	}
	return
}
