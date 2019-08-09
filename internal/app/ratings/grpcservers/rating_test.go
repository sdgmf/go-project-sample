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

var configFile = flag.String("f", "ratings.yml", "set config file which viper will loading.")

func TestRatingsServer_Get(t *testing.T) {
	flag.Parse()

	service := new(mocks.RatingsService)

	service.On("Get", mock.AnythingOfType("uint64")).Return(func(productID uint64) (p *models.Rating) {
		return &models.Rating{
			ProductID: productID,
		}
	}, func(ID uint64) error {
		return nil
	})

	server, err := CreateRatingsServer(*configFile, service)
	if err != nil {
		t.Fatalf("create product server error,%+v", err)
	}

	// 表格驱动测试
	tests := []struct {
		name     string
		id       uint64
		expected uint64
	}{
		{"id=1", 1, 1},
		{"id=2", 2, 2},
		{"id=3", 3, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := &proto.GetRatingRequest{
				ProductID: test.id,
			}
			r, err := server.Get(context.Background(), req)
			if err != nil {
				t.Fatalf("get detail error,%+v", err)
			}

			assert.Equal(t, test.expected, r.ProductID)
		})
	}

}
