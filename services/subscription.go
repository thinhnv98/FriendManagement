package services

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/repositories"
)

type ISubscriptionService interface {
	CreateSubscription(*model.SubscriptionServiceInput) error
	IsExistedSubscription(int, int) (bool, error)
	IsBlockedByOtherEmail(int, int) (bool, error)
}

type SubscriptionService struct {
	ISubscriptionRepo repositories.ISubscriptionRepo
}

func (_self SubscriptionService) CreateSubscription(subscriptionServiceInput *model.SubscriptionServiceInput) error {
	//Create repo input model
	repoInput := &model.SubscriptionRepoInput{
		Requestor: subscriptionServiceInput.Requestor,
		Target:    subscriptionServiceInput.Target,
	}
	err := _self.ISubscriptionRepo.CreateSubscription(repoInput)
	return err
}

func (_self SubscriptionService) IsExistedSubscription(requestorID int, targetID int) (bool, error) {
	exist, err := _self.ISubscriptionRepo.IsExistedSubscription(requestorID, targetID)
	return exist, err
}

func (_self SubscriptionService) IsBlockedByOtherEmail(requestorID int, targetID int) (bool, error) {
	blocked, err := _self.ISubscriptionRepo.IsBlockedByOtherEmail(requestorID, targetID)
	return blocked, err
}
