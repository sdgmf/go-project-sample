package grpcservers

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"github.com/sdgmf/go-project-sample/api/proto"
	"github.com/sdgmf/go-project-sample/internal/app/details/services"
	"go.uber.org/zap"
)

type DetailsServer struct {
	logger  *zap.Logger
	service services.DetailsService
}

func NewDetailsServer(logger *zap.Logger, ps services.DetailsService) (*DetailsServer, error) {
	return &DetailsServer{
		logger:  logger,
		service: ps,
	}, nil
}

func (s *DetailsServer) Get(ctx context.Context, req *proto.GetDetailRequest) (*proto.Detail, error) {
	p, err := s.service.Get(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "details grpc service get detail error")
	}
	ct, err := ptypes.TimestampProto(p.CreatedTime)
	if err != nil {
		return nil, errors.Wrap(err, "convert create time error")
	}

	resp := &proto.Detail{
		Id:          uint64(p.ID),
		Name:        p.Name,
		Price:       p.Price,
		CreatedTime: ct,
	}

	return resp, nil
}
