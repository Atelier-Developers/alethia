package impl

import (
	"context"
	"github.com/Atelier-Developers/alethia/interfaces"
	"github.com/Atelier-Developers/alethia/proto"
)

type BirthdayUpdaterServiceGrpcImpl struct {
	notificationHandler interfaces.NotificationHandler
	proto.UnimplementedBirthdayUpdaterServer
}

func (serviceImpl *BirthdayUpdaterServiceGrpcImpl) UpdateInfos(ctx context.Context, request *proto.BirthdayUpdateRequest) (*proto.BirthdayUpdateResponse, error) {

	err := serviceImpl.notificationHandler.UpdateBirthdayNotifications()
	if err != nil {
		return nil, err
	}

	return &proto.BirthdayUpdateResponse{Status: "Updated Successfully"}, nil
}

func NewBirthdayUpdaterServiceGrpcImpl(notificationHandler interfaces.NotificationHandler) *BirthdayUpdaterServiceGrpcImpl {
	return &BirthdayUpdaterServiceGrpcImpl{
		notificationHandler: notificationHandler,
	}
}
