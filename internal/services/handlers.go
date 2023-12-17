package services

import (
	"context"
	"fmt"

	pb "github.com/icrxz/grpc-go/pkg/pb/greeting/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreeterService interface {
	Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error)
}

type greeterService struct{}

func NewGreeterService() GreeterService {
	return &greeterService{}
}

func (gs *greeterService) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	if req.Msg == nil {
		err := status.New(codes.InvalidArgument, "Message cannot be empty").Err()
		return nil, err
	}

	helloMsg := fmt.Sprintf("%s, %s!", req.Msg.Greeting, req.Msg.Name)

	res := &pb.GreetResponse{
		Resp: helloMsg,
	}

	return res, nil
}
