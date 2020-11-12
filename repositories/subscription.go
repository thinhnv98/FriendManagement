package repositories

import (
	"database/sql"

	"S3_FriendManagement_ThinhNguyen/model"
)

type ISubscriptionRepo interface {
	CreateSubscription(*model.SubscriptionRepoInput) error
	IsExistedSubscription(int, int) (bool, error)
	IsBlockedByOtherEmail(int, int) (bool, error)
}

type SubscriptionRepo struct {
	Db *sql.DB
}

func (_self SubscriptionRepo) CreateSubscription(model *model.SubscriptionRepoInput) error {
	query := `insert into subscriptions(requestorid, targetid) VALUES ($1, $2)`
	_, err := _self.Db.Exec(query, model.Requestor, model.Target)
	return err
}

func (_self SubscriptionRepo) IsExistedSubscription(requestorID int, targetID int) (bool, error) {
	query := `select exists(select true from subscriptions where requestorid=$1 AND targetid=$2)`
	var exist bool
	err := _self.Db.QueryRow(query, requestorID, targetID).Scan(&exist)
	if err != nil {
		return true, err
	}
	if exist {
		return true, nil
	}
	return false, nil
}

func (_self SubscriptionRepo) IsBlockedByOtherEmail(requestorID int, targetID int) (bool, error) {
	query := `select exists(select true from blocks where (requestorid=$1 and targetid=$2) or (requestorid=$2 and targetid=$1))`
	var isBlock bool
	err := _self.Db.QueryRow(query, requestorID, targetID).Scan(&isBlock)
	if err != nil {
		return true, err
	}
	if isBlock {
		return true, nil
	}
	return false, nil
}
