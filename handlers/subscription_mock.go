package handlers

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"github.com/stretchr/testify/mock"
)

type mockSubscriptionService struct {
	mock.Mock
}

func (_self mockSubscriptionService) CreateSubscription(subscriptionServiceInput *model.SubscriptionServiceInput) error {
	args := _self.Called(subscriptionServiceInput)
	var r error
	if args.Get(0) != nil {
		r = args.Get(0).(error)
	}
	return r
}

func (_self mockSubscriptionService) IsExistedSubscription(requestorid int, targetid int) (bool, error) {
	args := _self.Called(requestorid, targetid)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockSubscriptionService) IsBlockedByOtherEmail(requestorid int, targetid int) (bool, error) {
	args := _self.Called(requestorid, targetid)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}
