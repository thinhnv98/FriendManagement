package services

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"github.com/stretchr/testify/mock"
)

type mockFriendRepo struct {
	mock.Mock
}

func (_self mockFriendRepo) CreateFriend(friendsRepoInput *model.FriendsRepoInput) error {
	args := _self.Called(friendsRepoInput)
	var r error
	if args.Get(0) != nil {
		r = args.Get(0).(error)
	}
	return r
}

func (_self mockFriendRepo) GetFriendListByID(userID int) ([]int, error) {
	args := _self.Called(userID)
	r0 := args.Get(0).([]int)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockFriendRepo) GetBlockedListByID(userID int) ([]int, error) {
	args := _self.Called(userID)
	r0 := args.Get(0).([]int)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockFriendRepo) GetBlockingListByID(userID int) ([]int, error) {
	args := _self.Called(userID)
	r0 := args.Get(0).([]int)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockFriendRepo) IsBlockedByOtherEmail(firstUserID int, secondUserID int) (bool, error) {
	args := _self.Called(firstUserID, secondUserID)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockFriendRepo) IsExistedFriend(firstUserID int, secondUserID int) (bool, error) {
	args := _self.Called(firstUserID, secondUserID)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockFriendRepo) GetSubscriberList(userID int) ([]int, error) {
	args := _self.Called(userID)
	r0 := args.Get(0).([]int)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockFriendRepo) GetEmailsFriendOrSubscribedWithNoBlocked(userID int) ([]int, error) {
	args := _self.Called(userID)
	r0 := args.Get(0).([]int)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}
