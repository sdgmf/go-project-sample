package controllers

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sdgmf/go-project-sample/internal/pkg/models"
	"github.com/sdgmf/go-project-sample/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

var r *gin.Engine
var configFile = flag.String("f", "reviews.yml", "set config file which viper will loading.")

func setup() {
	r = gin.New()
}

func TestReviewsController_Get(t *testing.T) {
	flag.Parse()
	setup()

	sto := new(mocks.ReviewsRepository)

	sto.On("Query", mock.AnythingOfType("uint64")).Return(func(ID uint64) (p []*models.Review) {
		return []*models.Review{&models.Review{
			ID: ID,
		}}
	}, func(ID uint64) error {
		return nil
	})

	c, err := CreateReviewsController(*configFile, sto)
	if err != nil {
		t.Fatalf("create reviews service error,%+v", err)
	}

	r.GET("/reviews", c.Query)

	tests := []struct {
		name     string
		id       uint64
		expected  int
	}{
		{"1", 1, 1},
		{"2", 2, 1},
		{"3", 3, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uri := fmt.Sprintf("/reviews?productID=%d", test.id)
			// 构造get请求
			req := httptest.NewRequest("GET", uri, nil)
			// 初始化响应
			w := httptest.NewRecorder()

			// 调用相应的controller接口
			r.ServeHTTP(w, req)

			// 提取响应
			result := w.Result()
			defer func() {
				_ = result.Body.Close()
			}()

			// 读取响应body
			body, _ := ioutil.ReadAll(result.Body)
			var rs []*models.Review
			err := json.Unmarshal(body, &rs)
			if err != nil {
				t.Errorf("unmarshal response body error:%v", err)
			}

			assert.Equal(t, test.expected, len(rs))
		})
	}

}
