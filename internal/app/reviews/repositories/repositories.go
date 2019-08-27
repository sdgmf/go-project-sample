package repositories

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewMysqlReviewsRepository)
//var MockProviderSet = wire.NewSet(wire.InterfaceValue(new(ReviewsRepository),new(MockReviewsRepository)))
