package consul

import (
	"github.com/google/wire"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Options struct {
	Addr string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)
	if err = v.UnmarshalKey("consul", o); err != nil {
		return nil, errors.Wrapf(err, "viper unmarshal consul options error")
	}

	return o, nil
}

func New(o *Options) (*consulApi.Client, error) {

	// initialize consul
	var (
		consulCli *consulApi.Client
		err       error
	)

	consulCli, err = consulApi.NewClient(&consulApi.Config{
		Address: o.Addr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "create consul client error")
	}

	return consulCli, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
