package services

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"github.com/stretchr/testify/mock"
)

type mockSubscriptionRepo struct {
	mock.Mock
}

func (_self mockSubscriptionRepo) CreateSubscription(model *model.SubscriptionRepoInput) error {
	args := _self.Called(model)
	var r error
	if args.Get(0) != nil {
		r = args.Get(0).(error)
	}
	return r
}

func (_self mockSubscriptionRepo) IsExistedSubscription(requestorID int, targetID int) (bool, error) {
	args := _self.Called(requestorID, targetID)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockSubscriptionRepo) IsBlockedByOtherEmail(requestorID int, targetID int) (bool, error) {
	args := _self.Called(requestorID, targetID)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}
