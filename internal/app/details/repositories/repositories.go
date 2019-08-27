package repositories

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewMysqlDetailsRepository)
//var MockProviderSet = wire.NewSet(wire.InterfaceValue(new(DetailsRepository),new(MockDetailsRepository)))
