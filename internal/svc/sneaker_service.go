package svc

import (
	"context"
	"errors"
	"fmt"

	"github.com/Israel-Ferreira/sneakers-crud-grpc/internal/data"
	"github.com/Israel-Ferreira/sneakers-crud-grpc/internal/models"
	pb "github.com/Israel-Ferreira/sneakers-crud-grpc/pkg/pb/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SneakerService struct {
	pb.SneakerServiceServer
	MongoClient *mongo.Client
}

func (sns *SneakerService) GetAll(req *pb.SneakerGetAllReq, stream pb.SneakerService_GetAllServer) error {

	ctx := context.TODO()
	cursor, err := sns.MongoClient.Database("sneakers_db").Collection("sneakers").Find(ctx, bson.M{})

	if err != nil {
		return err
	}

	for cursor.Next(ctx) {
		var sneakerModel models.SneakerModel

		if err := cursor.Decode(&sneakerModel); err != nil {
			return err
		}

		resp := sneakerModel.ToPbSneakerResp()

		stream.Send(resp)
	}

	return nil
}

func (sns *SneakerService) GetById(ctx context.Context, req *pb.SneakerGetByIdReq) (*pb.Sneaker, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		return nil, err
	}

	filterQuery := bson.M{"_id": id}

	result := sns.MongoClient.Database("sneakers_db").Collection("sneakers").FindOne(context.TODO(), filterQuery)

	if result.Err() != nil {
		return nil, result.Err()
	}

	var sneaker models.SneakerModel

	if err = result.Decode(&sneaker); err != nil {
		return nil, err
	}

	resp := sneaker.ToPbSneakerResp()

	return resp, nil
}

func (sns *SneakerService) DeleteById(ctx context.Context, req *pb.SneakerGetByIdReq) (*pb.DeleteSneakerMsg, error) {
	return nil, nil
}

func (sns *SneakerService) AddSneaker(ctx context.Context, req *pb.SneakerReq) (*pb.Sneaker, error) {
	sneaker := data.NewSneaker(req)

	fmt.Println(req)

	if sneaker.Model == "" {
		return nil, errors.New("error: sneaker model cannot be blank")
	}

	if sneaker.Manufacturer == "" {
		return nil, errors.New("error: manufacturer cannot be blank")
	}

	if sneaker.Colorway == "" {
		return nil, errors.New("error: colorway of sneaker cannot be blank")
	}

	if len(sneaker.AvailableSizes) == 0 {
		return nil, errors.New("error: available sizes cannot be empty")
	}

	sneakerModel := models.NewSneakerModel(*sneaker)

	if _, err := sns.MongoClient.Database("sneakers_db").Collection("sneakers").InsertOne(ctx, sneakerModel); err != nil {
		fmt.Println(err)
		return nil, err
	}

	resp := sneakerModel.ToPbSneakerResp()
	fmt.Println(resp)

	return resp, nil
}

func (sns *SneakerService) GetAvailableSizes(in *pb.SneakerGetByIdReq, stream pb.SneakerService_GetAvailableSizesServer) error {

	id, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return err
	}

	filterQuery := bson.M{"_id": id}

	result := sns.MongoClient.Database("sneakers_db").Collection("sneakers").FindOne(context.TODO(), filterQuery)

	if result.Err() != nil {
		return result.Err()
	}

	var sneaker models.SneakerModel

	if err := result.Decode(&sneaker); err != nil {
		return err
	}

	for _, size := range sneaker.AvailableSizes {
		stream.Send(&pb.AvailableSize{Size: size})
	}

	return nil
}

func (sns *SneakerService) UpdateSneaker(ctx context.Context, req *pb.SneakerReq) (*pb.UpdateSneakerMsg, error) {
	return nil, nil
}
