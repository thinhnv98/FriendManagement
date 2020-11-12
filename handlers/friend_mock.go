package handlers

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"github.com/stretchr/testify/mock"
)

type mockFriendService struct {
	mock.Mock
}

func (_self mockFriendService) CreateFriend(model *model.FriendsServiceInput) error {
	args := _self.Called(model)
	var r error
	if args.Get(0) != nil {
		r = args.Get(0).(error)
	}
	return r
}

func (_self mockFriendService) IsBlockedByOtherEmail(firstUserID int, secondUserID int) (bool, error) {
	args := _self.Called(firstUserID, secondUserID)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockFriendService) IsExistedFriend(firstUserID int, secondUserID int) (bool, error) {
	args := _self.Called(firstUserID, secondUserID)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockFriendService) GetFriendListByID(userID int) ([]string, error) {
	args := _self.Called(userID)
	r0 := args.Get(0).([]string)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockFriendService) GetCommonFriendListByID(userIDList []int) ([]string, error) {
	args := _self.Called(userIDList)
	r0 := args.Get(0).([]string)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockFriendService) GetEmailsReceiveUpdate(userID int, text string) ([]string, error) {
	args := _self.Called(userID, text)
	r0 := args.Get(0).([]string)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}
