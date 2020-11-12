package routes

import (
	"S3_FriendManagement_ThinhNguyen/handlers"
	"S3_FriendManagement_ThinhNguyen/repositories"
	"S3_FriendManagement_ThinhNguyen/services"
	"database/sql"
	"github.com/go-chi/chi"
	"net/http"
)

func CreateRoutes(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	//Routes for user
	r.Route("/user", func(r chi.Router) {
		UserHandler := handlers.UserHandler{
			IUserService: services.UserService{
				IUserRepo: repositories.UserRepo{
					Db: db,
				},
			},
		}
		r.MethodFunc(http.MethodPost, "/", UserHandler.CreateUser)
	})

	//Routes for Friend
	r.Route("/friend", func(r chi.Router) {
		FriendHandler := handlers.FriendHandler{
			IUserService: services.UserService{
				IUserRepo: repositories.UserRepo{
					Db: db,
				},
			},
			IFriendServices: services.FriendService{
				IFriendRepo: repositories.FriendRepo{
					Db: db,
				},
				IUserRepo: repositories.UserRepo{
					Db: db,
				},
			},
		}
		r.MethodFunc(http.MethodPost, "/", FriendHandler.CreateFriend)
		r.MethodFunc(http.MethodGet, "/friends", FriendHandler.GetFriendListByEmail)
		r.MethodFunc(http.MethodGet, "/common-friends", FriendHandler.GetCommonFriendListByEmails)
		r.MethodFunc(http.MethodGet, "/emails-receive-update", FriendHandler.GetEmailsReceiveUpdate)
	})
	//Routes for Subscription
	r.Route("/subscription", func(r chi.Router) {
		subscriptionHandler := handlers.SubscriptionHandler{
			IUserService: services.UserService{
				IUserRepo: repositories.UserRepo{
					Db: db,
				},
			},
			ISubscriptionService: services.SubscriptionService{
				ISubscriptionRepo: repositories.SubscriptionRepo{
					Db: db,
				},
			},
		}
		r.MethodFunc(http.MethodPost, "/", subscriptionHandler.CreateSubscription)
	})
	//Routes for Blocking
	r.Route("/block", func(r chi.Router) {
		blockHandler := handlers.BlockHandler{
			IUserService: services.UserService{
				IUserRepo: repositories.UserRepo{
					Db: db,
				},
			},
			IBlockingService: services.BlockingService{
				IBlockingRepo: repositories.BlockingRepo{
					Db: db,
				},
			},
		}
		r.MethodFunc(http.MethodPost, "/", blockHandler.CreateBlocking)
	})
	return r
}
