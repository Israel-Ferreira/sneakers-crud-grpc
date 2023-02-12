package models

import (
	"time"

	"github.com/Israel-Ferreira/sneakers-crud-grpc/internal/data"
	pb "github.com/Israel-Ferreira/sneakers-crud-grpc/pkg/pb/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SneakerModel struct {
	ID             primitive.ObjectID `bson:"_id"`
	Model          string             `bson:"sneaker_model"`
	Manufacturer   string             `bson:"manufacturer"`
	IsDeprecated   bool               `bson:"is_deprecated"`
	Colorway       string             `bson:"colorway"`
	Price          float64            `bson:"price"`
	AvailableSizes []int32            `bson:"available_sizes"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
}

func (sneakerModel *SneakerModel) ToPbSneakerResp() *pb.Sneaker {
	return &pb.Sneaker{
		Id:             sneakerModel.ID.Hex(),
		Model:          sneakerModel.Model,
		Manufacturer:   sneakerModel.Manufacturer,
		Colorway:       sneakerModel.Colorway,
		Price:          sneakerModel.Price,
		Deprecated:     sneakerModel.IsDeprecated,
	}
}

func NewSneakerModel(sneaker data.Sneaker) *SneakerModel {
	return &SneakerModel{
		ID:             primitive.NewObjectID(),
		Model:          sneaker.Model,
		Manufacturer:   sneaker.Manufacturer,
		Colorway:       sneaker.Colorway,
		Price:          sneaker.Price,
		IsDeprecated:   sneaker.Deprecated,
		AvailableSizes: sneaker.AvailableSizes,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
