package services

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"github.com/stretchr/testify/mock"
)

type mockUserRepo struct {
	mock.Mock
}

func (_self mockUserRepo) CreateUser(userRepoInput *model.UserRepoInput) error {
	args := _self.Called(userRepoInput)
	var r error
	if args.Get(0) != nil {
		r = args.Get(0).(error)
	}
	return r
}

func (_self mockUserRepo) GetUserIDByEmail(email string) (int, error) {
	args := _self.Called(email)
	r0 := args.Get(0).(int)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserRepo) IsExistedUser(email string) (bool, error) {
	args := _self.Called(email)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserRepo) GetEmailListByIDs(userIDs []int) ([]string, error) {
	args := _self.Called(userIDs)
	r0 := args.Get(0).([]string)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserRepo) GetUserIDsByEmails(emails []string) ([]int, error) {
	args := _self.Called(emails)
	r0 := args.Get(0).([]int)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserRepo) CheckInvalidEmails(emails []string) ([]string, error) {
	args := _self.Called(emails)
	r0 := args.Get(0).([]string)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}
