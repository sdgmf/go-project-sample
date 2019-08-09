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
var configFile = flag.String("f", "ratings.yml", "set config file which viper will loading.")

func setup() {
	r = gin.New()
}

func TestRatingsController_Get(t *testing.T) {
	flag.Parse()
	setup()

	sto := new(mocks.RatingsRepository)

	sto.On("Get", mock.AnythingOfType("uint64")).Return(func(productID uint64) (p *models.Rating) {
		return &models.Rating{
			ProductID: productID,
		}
	}, func(ID uint64) error {
		return nil
	})

	c, err := CreateRatingsController(*configFile, sto)
	if err != nil {
		t.Fatalf("create product serviceerror,%+v", err)
	}

	r.GET("/rating/:productID", c.Get)

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
			uri := fmt.Sprintf("/rating/%d", test.id)
			// 构造get请求
			req := httptest.NewRequest("GET", uri, nil)
			// 初始化响应
			w := httptest.NewRecorder()

			// 调用相应的controller接口
			r.ServeHTTP(w, req)

			// 提取响应
			rs := w.Result()
			defer func() {
				_ = rs.Body.Close()
			}()

			// 读取响应body
			body, _ := ioutil.ReadAll(rs.Body)
			r := new(models.Rating)
			err := json.Unmarshal(body, r)
			if err != nil {
				t.Errorf("unmarshal response body error:%v", err)
			}

			assert.Equal(t, test.expected, r.ProductID)
		})
	}

}
