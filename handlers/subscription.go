package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/services"
)

type SubscriptionHandler struct {
	IUserService         services.IUserService
	ISubscriptionService services.ISubscriptionService
}

func (_self SubscriptionHandler) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	//Decode request body
	subscriptionRequest := model.CreateSubscriptionRequest{}
	if err := json.NewDecoder(r.Body).Decode(&subscriptionRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Validate request
	if err := subscriptionRequest.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Validate and get UserID by email
	userIDList, statusCode, err := _self.CreateSubscribeValidation(subscriptionRequest)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	//Create input services model
	modelServiceInput := &model.SubscriptionServiceInput{
		Requestor: userIDList[0],
		Target:    userIDList[1],
	}
	//Call services
	if err := _self.ISubscriptionService.CreateSubscription(modelServiceInput); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Success: true,
	})
	return
}

func (_self SubscriptionHandler) CreateSubscribeValidation(subscriptionRequest model.CreateSubscriptionRequest) ([]int, int, error) {
	//Check requestor email
	requestorUSerID, err := _self.IUserService.GetUserIDByEmail(subscriptionRequest.Requestor)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if requestorUSerID == 0 {
		return nil, http.StatusBadRequest, errors.New("requestor email does not exist")
	}

	//Check target email
	targetUserID, err := _self.IUserService.GetUserIDByEmail(subscriptionRequest.Target)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if targetUserID == 0 {
		return nil, http.StatusBadRequest, errors.New("target email does not exist")
	}

	//Check subscription existed
	exist, err := _self.ISubscriptionService.IsExistedSubscription(requestorUSerID, targetUserID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if exist {
		return nil, http.StatusAlreadyReported, errors.New("those email address have already subscribed the each other")
	}

	//Check blocked
	blocked, err := _self.ISubscriptionService.IsBlockedByOtherEmail(requestorUSerID, targetUserID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if blocked {
		return nil, http.StatusPreconditionFailed, errors.New("those emails have already been blocked by the each other")
	}
	return []int{requestorUSerID, targetUserID}, 0, nil
}
