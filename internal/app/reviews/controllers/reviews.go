package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"github.com/sdgmf/go-project-sample/internal/app/reviews/services"
)

type ReviewsController struct {
	logger  *zap.Logger
	service services.ReviewsService
}

func NewReviewsController(logger *zap.Logger, s services.ReviewsService) *ReviewsController {
	return &ReviewsController{
		logger:  logger,
		service: s,
	}
}

func (pc *ReviewsController) Query(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Query("productID"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	rs, err := pc.service.Query(ID)
	if err != nil {
		pc.logger.Error("get review by productID error", zap.Error(err))
		c.String(http.StatusInternalServerError,"%+v", err)
		return
	}

	c.JSON(http.StatusOK, rs)
}
