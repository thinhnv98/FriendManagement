package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/services"
)

type BlockHandler struct {
	IUserService     services.IUserService
	IBlockingService services.IBlockingService
}

func (_self BlockHandler) CreateBlocking(w http.ResponseWriter, r *http.Request) {
	//Decode request body
	blockingRequest := model.BlockingRequest{}
	if err := json.NewDecoder(r.Body).Decode(&blockingRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate request
	if err := blockingRequest.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validate and get UserID by email
	userIDList, statusCode, err := _self.createBlockingValidation(blockingRequest)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	//Create block services input model
	blockingServiceInput := &model.BlockingServiceInput{
		Requestor: userIDList[0],
		Target:    userIDList[1],
	}

	//Call services
	if err := _self.IBlockingService.CreateBlocking(blockingServiceInput); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Response
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Success: true,
	})
	return
}

func (_self BlockHandler) createBlockingValidation(blockingRequest model.BlockingRequest) ([]int, int, error) {
	// Get user id of the requestor
	requestorUserID, err := _self.IUserService.GetUserIDByEmail(blockingRequest.Requestor)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if requestorUserID == 0 {
		return nil, http.StatusBadRequest, errors.New("the requestor does not exist")
	}

	// Get user id of the target
	targetUserID, err := _self.IUserService.GetUserIDByEmail(blockingRequest.Target)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if targetUserID == 0 {
		return nil, http.StatusBadRequest, errors.New("the target does not exist")
	}

	//Check blocked
	blocked, err := _self.IBlockingService.IsExistedBlocking(requestorUserID, targetUserID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if blocked {
		return nil, http.StatusPreconditionFailed, errors.New("target's email have already been blocked by requestor's email")
	}
	return []int{requestorUserID, targetUserID}, 0, nil
}
