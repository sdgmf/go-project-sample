package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sdgmf/go-project-sample/internal/pkg/transports/http"
)

func CreateInitControllersFn(
	pc *ReviewsController,
) http.InitControllers {
	return func(r *gin.Engine) {
		r.GET("/reviews", pc.Query)
	}
}

var ProviderSet = wire.NewSet(NewReviewsController, CreateInitControllersFn)
