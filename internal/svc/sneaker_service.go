package svc

import (
	"context"

	"github.com/Israel-Ferreira/sneakers-crud-grpc/internal/data"
	pb "github.com/Israel-Ferreira/sneakers-crud-grpc/pkg/pb/proto"
)

type SneakerService struct {
	pb.SneakerServiceServer
}

func (sns *SneakerService) GetAll(req *pb.SneakerGetAllReq, stream pb.SneakerService_GetAllServer) error {

	for _, snkr := range data.Sneakers {
		stream.Send(snkr.SneakerToPbModel())
	}

	return nil
}

func (sns *SneakerService) GetById(ctx context.Context, req *pb.SneakerGetByIdReq) (*pb.Sneaker, error) {
	return nil, nil
}

func (sns *SneakerService) DeleteById(ctx context.Context, req *pb.SneakerGetByIdReq) (*pb.DeleteSneakerMsg, error) {
	return nil, nil
}

func (sns *SneakerService) AddSneaker(ctx context.Context, req *pb.SneakerReq) (*pb.Sneaker, error) {
	return nil, nil
}

func (sns *SneakerService) UpdateSneaker(ctx context.Context, req *pb.SneakerReq) (*pb.UpdateSneakerMsg, error) {
	return nil, nil
}
