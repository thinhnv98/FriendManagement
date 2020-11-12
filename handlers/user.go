package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/services"
)

type UserHandler struct {
	IUserService services.IUserService
}

func (_self *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	//Decode request body
	userRequest := model.UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Validation
	if err := userRequest.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if statusCode, err := _self.IsExistedUser(userRequest.Email); err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	//Convert to services input model
	userServiceInp := &model.UserServiceInput{
		Email: userRequest.Email,
	}

	//Call services
	if err := _self.IUserService.CreateUser(userServiceInp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Response
	json.NewEncoder(w).Encode(&model.SuccessResponse{
		Success: true,
	})
}

func (_self *UserHandler) IsExistedUser(email string) (int, error) {
	//Call services
	existed, err := _self.IUserService.IsExistedUser(email)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if existed {
		return http.StatusAlreadyReported, errors.New("this email address existed")
	}
	return 0, nil
}
