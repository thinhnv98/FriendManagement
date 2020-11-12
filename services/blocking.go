package services

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/repositories"
)

type IBlockingService interface {
	CreateBlocking(*model.BlockingServiceInput) error
	IsExistedBlocking(int, int) (bool, error)
}

type BlockingService struct {
	IBlockingRepo repositories.IBlockingRepo
}

func (_self BlockingService) CreateBlocking(blocking *model.BlockingServiceInput) error {
	//Create repo input model
	blockingRepoInputModel := &model.BlockingRepoInput{
		Requestor: blocking.Requestor,
		Target:    blocking.Target,
	}
	err := _self.IBlockingRepo.CreateBlocking(blockingRepoInputModel)
	return err
}

func (_self BlockingService) IsExistedBlocking(requestorID int, targetID int) (bool, error) {
	exist, err := _self.IBlockingRepo.IsExistedBlocking(requestorID, targetID)
	return exist, err
}
