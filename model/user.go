package model

import (
	"errors"

	"S3_FriendManagement_ThinhNguyen/utils"
)

type User struct {
	Email string
}

//model handler
type UserRequest struct {
	Email string `json:"email"`
}

func (_self UserRequest) Validate() error {
	if _self.Email == "" {
		return errors.New("\"email\" is required")
	}

	isValid, err := utils.IsValidEmail(_self.Email)
	if err != nil {
		return errors.New("validate \"email\" format failed")
	}
	if !isValid {
		return errors.New("\"email\"'s format is not valid. (ex: \"andy@abc.xyz\")")
	}
	return nil
}

type SuccessResponse struct {
	Success bool `json:"Success"`
}

//model services
type UserServiceInput struct {
	Email string `json:"email"`
}

//model repo
type UserRepoInput struct {
	Email string `json:"email"`
}
