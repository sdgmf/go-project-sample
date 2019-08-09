package services

import (
	"flag"
	"github.com/sdgmf/go-project-sample/internal/pkg/models"
	"github.com/sdgmf/go-project-sample/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var configFile = flag.String("f", "ratings.yml", "set config file which viper will loading.")

func TestDefaultRatingsService_Get(t *testing.T) {
	flag.Parse()

	sto := new(mocks.RatingsRepository)

	sto.On("Get", mock.AnythingOfType("uint64")).Return(func(ID uint64) (p *models.Rating) {
		return &models.Rating{
			ID: ID,
		}
	}, func(ID uint64) error {
		return nil
	})

	svc, err := CreateRatingsService(*configFile, sto)
	if err != nil {
		t.Fatalf("create product serviceerror,%+v", err)
	}

	// 表格驱动测试
	tests := []struct {
		name     string
		id       uint64
		expected uint64
	}{
		{"1+1", 1, 1},
		{"2+3", 2, 2},
		{"4+5", 3, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p, err := svc.Get(test.id)
			if err != nil {
				t.Fatalf("product service get proudct error,%+v", err)
			}

			assert.Equal(t, test.expected, p.ID)
		})
	}
}
