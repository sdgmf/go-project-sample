package services

import (
	"flag"
	"github.com/sdgmf/go-project-sample/internal/pkg/models"
	"github.com/sdgmf/go-project-sample/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var configFile = flag.String("f", "reviews.yml", "set config file which viper will loading.")

func TestReviewsService_Query(t *testing.T) {
	flag.Parse()

	sto := new(mocks.ReviewsRepository)

	sto.On("Query", mock.AnythingOfType("uint64")).Return(func(productID uint64) (p []*models.Review) {
		return []*models.Review{&models.Review{
			ProductID: productID,
		}}
	}, func(ID uint64) error {
		return nil
	})

	svc, err := CreateReviewsService(*configFile, sto)
	if err != nil {
		t.Fatalf("create reviews service error,%+v", err)
	}

	// 表格驱动测试
	tests := []struct {
		name     string
		id       uint64
		expected int
	}{
		{"id=1", 1, 1},
		{"id=2", 2, 1},
		{"id=3", 3, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rs, err := svc.Query(test.id)
			if err != nil {
				t.Fatalf("product service get proudct error,%+v", err)
			}

			assert.Equal(t, test.expected, len(rs))
		})
	}
}
