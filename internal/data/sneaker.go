package data

import (
	pb "github.com/Israel-Ferreira/sneakers-crud-grpc/pkg/pb/proto"
)

type Sneaker struct {
	Id             string
	Model          string
	Manufacturer   string
	Colorway       string
	Price          float64
	Deprecated     bool
	AvailableSizes []int32
}

func (snkr *Sneaker) SneakerToPbModel() *pb.Sneaker {
	return &pb.Sneaker{
		Id:             snkr.Id,
		Model:          snkr.Model,
		Colorway:       snkr.Colorway,
		Manufacturer:   snkr.Manufacturer,
		Price:          snkr.Price,
		Deprecated:     snkr.Deprecated,
	}
}

func NewSneaker(pbSneaker *pb.SneakerReq) *Sneaker {
	return &Sneaker{
		Model:          pbSneaker.Model,
		Manufacturer:   pbSneaker.Manufacturer,
		Colorway:       pbSneaker.Colorway,
		Price:          pbSneaker.Price,
		Deprecated:     pbSneaker.Deprecated,
		AvailableSizes: pbSneaker.AvailableSizes,
	}
}

var Sneakers = []Sneaker{
	{
		Id:             "1",
		Model:          "AIR FORCE 1 LOW",
		Manufacturer:   "NIKE",
		Colorway:       "WHITE",
		Price:          799.90,
		AvailableSizes: []int32{37, 38, 39, 40, 41, 42, 43, 44, 45, 46},
	},

	{
		Id:             "2",
		Model:          "FORUM HIGH",
		Manufacturer:   "ADIDAS",
		Colorway:       "WHITE/BLACK",
		Price:          699.90,
		AvailableSizes: []int32{37, 38, 39, 40, 41, 42, 43, 44, 45, 46},
	},

	{
		Id:             "3",
		Model:          "CHUCK TAYLOR ALL STAR",
		Manufacturer:   "CONVERSE",
		Colorway:       "TRADICIONAL",
		Price:          200.00,
		AvailableSizes: []int32{37, 38, 39, 40, 41, 42, 43, 44, 45, 46},
	},
	{
		Id:             "4",
		Model:          "SUEDE CLASSIC",
		Manufacturer:   "PUMA",
		Colorway:       "PRETO",
		Price:          499.90,
		AvailableSizes: []int32{37, 38, 39, 40, 41, 42, 43, 44, 45, 46},
	},
}
