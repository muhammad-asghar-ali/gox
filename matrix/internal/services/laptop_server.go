package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"matrix/internal/pb"
)

type (
	LaptopServer struct {
		Store LaptopStore
	}
)

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{Store: store}
}

func (s *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	lp := req.GetLaptop()

	if len(lp.Id) > 0 {
		_, err := uuid.Parse(lp.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid UUID format: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		lp.Id = id.String()
	}

	err := s.Store.Save(lp)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}

		return nil, status.Error(code, err.Error())
	}

	res := &pb.CreateLaptopResponse{
		Id: lp.Id,
	}

	return res, nil
}
