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

var configFile = flag.String("f", "details.yml", "set config file which viper will loading.")

func TestDetailsService_Get(t *testing.T) {
	flag.Parse()

	service := new(mocks.DetailsService)

	service.On("Get", mock.AnythingOfType("uint64")).Return(func(ID uint64) (p *models.Detail) {
		return &models.Detail{
			ID: ID,
		}
	}, func(ID uint64) error {
		return nil
	})

	server, err := CreateDetailsServer(*configFile, service)
	if err != nil {
		t.Fatalf("create product server error,%+v", err)
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
			req := &proto.GetDetailRequest{
				Id: test.id,
			}
			p, err := server.Get(context.Background(), req)
			if err != nil {
				t.Fatalf("product service get proudct error,%+v", err)
			}

			assert.Equal(t, test.expected, p.Id)
		})
	}

}
