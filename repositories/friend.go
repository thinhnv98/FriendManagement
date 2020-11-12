package repositories

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"database/sql"
)

type IFriendRepo interface {
	CreateFriend(*model.FriendsRepoInput) error
	GetFriendListByID(int) ([]int, error)
	GetBlockedListByID(int) ([]int, error)
	GetBlockingListByID(int) ([]int, error)
	IsBlockedByOtherEmail(int, int) (bool, error)
	IsExistedFriend(int, int) (bool, error)
	GetSubscriberList(int) ([]int, error)
	GetEmailsFriendOrSubscribedWithNoBlocked(int) ([]int, error)
}

type FriendRepo struct {
	Db *sql.DB
}

func (_self FriendRepo) CreateFriend(friendsRepoInput *model.FriendsRepoInput) error {
	query := `insert into friends(firstid, secondid) values ($1, $2)`
	_, err := _self.Db.Exec(query, friendsRepoInput.FirstID, friendsRepoInput.SecondID)
	return err
}

func (_self FriendRepo) GetFriendListByID(userID int) ([]int, error) {
	query := `select firstid, secondid from friends where firstid=$1 or secondid = $1`

	var friendListID = make([]int, 0)
	rows, err := _self.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var firstID, secondID int
		if err := rows.Scan(&firstID, &secondID); err != nil {
			return nil, err
		}
		if firstID == userID {
			friendListID = append(friendListID, secondID)
		}
		if secondID == userID {
			friendListID = append(friendListID, firstID)
		}
	}
	return friendListID, err
}

func (_self FriendRepo) GetBlockingListByID(userID int) ([]int, error) {
	query := `select targetid from blocks where requestorid = $1`

	var blockedListID = make([]int, 0)
	rows, err := _self.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var blockedUserID int
		if err := rows.Scan(&blockedUserID); err != nil {
			return nil, err
		}
		blockedListID = append(blockedListID, blockedUserID)
	}
	return blockedListID, err
}

func (_self FriendRepo) GetBlockedListByID(userID int) ([]int, error) {
	query := `select requestorid from blocks where targetid = $1`

	var blockingListID = make([]int, 0)
	rows, err := _self.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var blockingUserID int
		if err := rows.Scan(&blockingUserID); err != nil {
			return nil, err
		}
		blockingListID = append(blockingListID, blockingUserID)
	}
	return blockingListID, err
}

func (_self FriendRepo) IsBlockedByOtherEmail(firstUserID int, secondUserID int) (bool, error) {
	query := `select exists(select true from blocks WHERE (
    						    	requestorid in ($1, $2) 
								    AND 
    						    	targetid in ($1, $2)
    						      ))`
	var isBlocked bool
	err := _self.Db.QueryRow(query, firstUserID, secondUserID).Scan(&isBlocked)
	if err != nil {
		return true, err
	}
	if isBlocked {
		return true, nil
	}
	return false, nil
}

func (_self FriendRepo) IsExistedFriend(firstUserID int, secondUserID int) (bool, error) {
	query := `select exists(
    						select true 
    						from friends 
    						where (
    						    	firstid in ($1, $2) 
								    AND 
    						    	secondid in ($1, $2)
    						      )
    						)`
	var existed bool
	err := _self.Db.QueryRow(query, firstUserID, secondUserID).Scan(&existed)
	if err != nil {
		return true, err
	}
	if existed {
		return true, nil
	}
	return false, nil
}

func (_self FriendRepo) GetSubscriberList(userID int) ([]int, error) {
	query := `select requestorid from subscriptions where targetid=$1`
	subscribers := make([]int, 0)
	rows, err := _self.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		subscribers = append(subscribers, id)
	}
	return subscribers, nil
}

func (_self FriendRepo) GetEmailsFriendOrSubscribedWithNoBlocked(userID int) ([]int, error) {
	query := `select distinct val.ID
			  from
				(
					select ue.ID
					from
						useremails ue
							join friends f
								 on (ue.id = f.firstid or ue.id = f.secondid)
					where ue.id <> $1
					  and (f.firstid = $1 or f.secondid = $1)
					union
					select ue.ID
					from
						subscriptions s
							join useremails ue
								 on s.targetid = ue.id
					where ue.id <> $1
				) as val
				where not exists(
				    select 1
				    from blocks b 
				    where b.requestorid = val.id
				    and b.targetid = $1
				)`
	rows, err := _self.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	UserIDs := make([]int, 0)
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		UserIDs = append(UserIDs, id)
	}
	return UserIDs, nil
}
