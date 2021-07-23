package impl

import (
	"context"
	"github.com/Atelier-Developers/alethia/proto"
	"log"
)

type BirthdayUpdaterServiceGrpcImpl struct {
	proto.UnimplementedBirthdayUpdaterServer
}

func (serviceImpl *BirthdayUpdaterServiceGrpcImpl) UpdateInfos(ctx context.Context, request *proto.BirthdayUpdateRequest) (*proto.BirthdayUpdateResponse, error) {

	log.Println("Repository persisted to the storage")

	return &proto.BirthdayUpdateResponse{Status: "Well done"}, nil
}

func NewBirthdayUpdaterServiceGrpcImpl() *BirthdayUpdaterServiceGrpcImpl {
	return &BirthdayUpdaterServiceGrpcImpl{}
}
