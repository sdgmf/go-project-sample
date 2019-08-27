package repositories

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewMysqlRatingsRepository)
//var MockProviderSet = wire.NewSet(wire.InterfaceValue(new(RatingsRepository),new(MockRatingsRepository)))
