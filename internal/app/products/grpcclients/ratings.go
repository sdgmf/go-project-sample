package grpcclients

import (
	"github.com/pkg/errors"
	"github.com/sdgmf/go-project-sample/api/proto"
	"github.com/sdgmf/go-project-sample/internal/pkg/transports/grpc"
)

func NewRatingsClient(client *grpc.Client) (proto.RatingsClient, error) {
	conn, err := client.Dial("Ratings")
	if err != nil {
		return nil, errors.Wrap(err, "detail client dial error")
	}
	c := proto.NewRatingsClient(conn)

	return c, nil
}
