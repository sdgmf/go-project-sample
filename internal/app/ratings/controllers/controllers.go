package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sdgmf/go-project-sample/internal/pkg/transports/http"
)

func CreateInitControllersFn(
	pc *RatingsController,
) http.InitControllers {
	return func(r *gin.Engine) {
		r.GET("/rating/:productID", pc.Get)
	}
}

var ProviderSet = wire.NewSet(NewRatingsController, CreateInitControllersFn)
