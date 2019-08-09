package services

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"github.com/sdgmf/go-project-sample/api/proto"
	"github.com/sdgmf/go-project-sample/internal/pkg/models"
	"go.uber.org/zap"
)

type ProductsService interface {
	Get(c context.Context, ID uint64) (*models.Product, error)
}

type DefaultProductsService struct {
	logger     *zap.Logger
	detailsSvc proto.DetailsClient
	ratingsSvc proto.RatingsClient
	reviewsSvc proto.ReviewsClient
}

func NewProductService(logger *zap.Logger,
	detailsSvc proto.DetailsClient,
	ratingsSvc proto.RatingsClient,
	reviewsSvc proto.ReviewsClient) ProductsService {
	return &DefaultProductsService{
		logger:     logger.With(zap.String("type", "DefaultProductsService")),
		detailsSvc: detailsSvc,
		ratingsSvc: ratingsSvc,
		reviewsSvc: reviewsSvc,
	}
}

func (s *DefaultProductsService) Get(c context.Context, productID uint64) (p *models.Product, err error) {
	var (
		detail  *models.Detail
		rating  *models.Rating
		reviews []*models.Review
	)

	// get detail
	{
		req := &proto.GetDetailRequest{
			Id: productID,
		}

		pd, err := s.detailsSvc.Get(c, req)
		if err != nil {
			return nil, errors.Wrap(err, "get rating error")
		}
		ct, err := ptypes.Timestamp(pd.CreatedTime)

		detail = &models.Detail{
			ID:          pd.Id,
			Price:       pd.Price,
			CreatedTime: ct,
		}
	}

	// get rating
	{
		req := &proto.GetRatingRequest{
			ProductID: productID,
		}

		pr, err := s.ratingsSvc.Get(c, req)
		if err != nil {
			return nil, errors.Wrap(err, "get rating error")
		}
		ut, err := ptypes.Timestamp(pr.UpdatedTime)

		rating = &models.Rating{
			ID:          pr.Id,
			Score:       pr.Score,
			UpdatedTime: ut,
		}
	}

	// query reviews
	{
		req := &proto.QueryReviewsRequest{
			ProductID: productID,
		}

		resp, err := s.reviewsSvc.Query(c, req)
		if err != nil {
			return nil, errors.Wrap(err, "get rating error")
		}

		reviews = make([]*models.Review, 0, len(resp.Reviews))

		for _, pr := range resp.Reviews {
			ct, err := ptypes.Timestamp(pr.CreatedTime)
			if err != nil {
				return nil, errors.Wrapf(err, "convert create time error")
			}

			r := &models.Review{
				ID:          pr.Id,
				ProductID:   pr.ProductID,
				Message:     pr.Message,
				CreatedTime: ct,
			}

			reviews = append(reviews, r)
		}

	}

	return &models.Product{
		Detail:  detail,
		Rating:  rating,
		Reviews: reviews,
	}, nil
}
