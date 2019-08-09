package grpcservers

import (
	"context"
	"flag"
	"github.com/sdgmf/go-project-sample/api/proto"
	"github.com/sdgmf/go-project-sample/internal/pkg/models"
	"github.com/sdgmf/go-project-sample/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var configFile = flag.String("f", "reviews.yml", "set config file which viper will loading.")

func TestReviewsServer_Query(t *testing.T) {
	flag.Parse()

	service := new(mocks.ReviewsService)

	service.On("Query", mock.AnythingOfType("uint64")).Return(func(proudctID uint64) (p []*models.Review) {
		return []*models.Review{&models.Review{
			ProductID: proudctID,
		}}
	}, func(ID uint64) error {
		return nil
	})

	server, err := CreateReviewsServer(*configFile, service)
	if err != nil {
		t.Fatalf("create product server error,%+v", err)
	}

	// 表格驱动测试
	tests := []struct {
		name     string
		id       uint64
		expected int
	}{
		{"productID=1", 1, 1},
		{"productID=2", 2, 1},
		{"productID=3", 3, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := &proto.QueryReviewsRequest{
				ProductID: test.id,
			}
			rs, err := server.Query(context.Background(), req)
			if err != nil {
				t.Fatalf("product service get proudct error,%+v", err)
			}

			assert.Equal(t, test.expected, len(rs.Reviews))
		})
	}

}
