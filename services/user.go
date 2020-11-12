package services

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/repositories"
)

type IUserService interface {
	CreateUser(*model.UserServiceInput) error
	IsExistedUser(string) (bool, error)
	GetUserIDByEmail(string) (int, error)
	CheckInvalidEmails([]string) ([]string, error)
}

type UserService struct {
	IUserRepo repositories.IUserRepo
}

func (_self UserService) CreateUser(userServiceInput *model.UserServiceInput) error {
	//Convert to repo input
	userRepoInput := &model.UserRepoInput{
		Email: userServiceInput.Email,
	}

	err := _self.IUserRepo.CreateUser(userRepoInput)
	return err
}

func (_self UserService) GetUserIDByEmail(email string) (int, error) {
	result, err := _self.IUserRepo.GetUserIDByEmail(email)
	return result, err
}

func (_self UserService) IsExistedUser(email string) (bool, error) {
	//call repo
	existed, err := _self.IUserRepo.IsExistedUser(email)
	return existed, err
}

func (_self UserService) CheckInvalidEmails(emails []string) ([]string, error) {
	results, err := _self.IUserRepo.CheckInvalidEmails(emails)
	return results, err
}
