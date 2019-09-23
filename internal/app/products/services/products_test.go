package services

import (
	"context"
	"flag"
	"github.com/golang/protobuf/ptypes"
	"github.com/sdgmf/go-project-sample/api/proto"
	"github.com/sdgmf/go-project-sample/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"testing"
)

var configFile = flag.String("f", "products.bak", "set config file which viper will loading.")

func TestDefaultProductsService_Get(t *testing.T) {
	flag.Parse()

	detailsCli := new(mocks.DetailsClient)
	detailsCli.On("Get", mock.Anything, mock.Anything).
		Return(func(ctx context.Context, req *proto.GetDetailRequest, cos ...grpc.CallOption) *proto.Detail {
			return &proto.Detail{
				Id:          req.Id,
				CreatedTime: ptypes.TimestampNow(),
			}
		}, func(ctx context.Context, req *proto.GetDetailRequest, cos ...grpc.CallOption) error {
			return nil
		})

	ratingsCli := new(mocks.RatingsClient)

	ratingsCli.On("Get", context.Background(), mock.AnythingOfType("*proto.GetRatingRequest")).
		Return(func(ctx context.Context, req *proto.GetRatingRequest, cos ...grpc.CallOption) *proto.Rating {
			return &proto.Rating{
				Id:          req.ProductID,
				UpdatedTime: ptypes.TimestampNow(),
			}
		}, func(ctx context.Context, req *proto.GetRatingRequest, cos ...grpc.CallOption) error {
			return nil
		})

	reviewsCli := new(mocks.ReviewsClient)

	reviewsCli.On("Query", context.Background(), mock.AnythingOfType("*proto.QueryReviewsRequest")).
		Return(func(ctx context.Context, req *proto.QueryReviewsRequest, cos ...grpc.CallOption) *proto.QueryReviewsResponse {
			return &proto.QueryReviewsResponse{
				Reviews: []*proto.Review{
					&proto.Review{
						Id:          req.ProductID,
						CreatedTime: ptypes.TimestampNow(),
					},
				},
			}
		}, func(ctx context.Context, req *proto.QueryReviewsRequest, cos ...grpc.CallOption) error {
			return nil
		})

	svc, err := CreateProductsService(*configFile, detailsCli, ratingsCli, reviewsCli)
	if err != nil {
		t.Fatalf("create product service error,%+v", err)
	}

	// 表格驱动测试
	tests := []struct {
		name     string
		id       uint64
		expected bool
	}{
		{"id=1", 1, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := svc.Get(context.Background(), test.id)

			if test.expected {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
