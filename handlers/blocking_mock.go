package handlers

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"github.com/stretchr/testify/mock"
)

type mockBlockingService struct {
	mock.Mock
}

func (_self mockBlockingService) CreateBlocking(input *model.BlockingServiceInput) error {
	args := _self.Called(input)
	var r error
	if args.Get(0) != nil {
		r = args.Get(0).(error)
	}
	return r
}

func (_self mockBlockingService) IsExistedBlocking(requestorID int, targetID int) (bool, error) {
	args := _self.Called(requestorID, targetID)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}
