package routes

import (
	api "eaas_back/app/handlers"
	submissionRepo "eaas_back/app/repositories/submission"
	userRepo "eaas_back/app/repositories/user"
	authSrv "eaas_back/app/services/auth"
	submissionSrv "eaas_back/app/services/submission"
	userSrv "eaas_back/app/services/user"
	"eaas_back/config"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

var (
	BaseRoute = "/api/v1"
)

func InitializeRoutes(router *mux.Router, dbSession *mgo.Session, conf *config.Configuration) {
	userRepository := userRepo.New(dbSession, conf)
	userService := userSrv.New(userRepository)
	userAPI := api.NewUserAPI(userService)

	submissionRepository := submissionRepo.New(dbSession, conf)
	submissionService := submissionSrv.New(submissionRepository)
	submissionAPI := api.NewSubmissionAPI(submissionService)

	authService := authSrv.New(userRepository)
	authAPI := api.NewAuthAPI(authService, conf)

	// Routes

	//  -------------------------- Auth APIs ------------------------------------
	router.HandleFunc(BaseRoute+"/auth/register", authAPI.Create).Methods(http.MethodPost)
	router.HandleFunc(BaseRoute+"/auth/login", authAPI.Login).Methods(http.MethodPost)

	// -------------------------- User APIs ------------------------------------
	router.HandleFunc(BaseRoute+"/users/me", userAPI.Get).Methods(http.MethodGet)
	router.HandleFunc(BaseRoute+"/users", userAPI.Update).Methods(http.MethodPut)

	//  -------------------------- Submission APIs ------------------------------
	router.HandleFunc(BaseRoute+"/submission", submissionAPI.Get).Methods(http.MethodGet)
	router.HandleFunc(BaseRoute+"/submission", submissionAPI.Create).Methods(http.MethodPost)
}
