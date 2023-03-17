package routes

import (
	"project/controller"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	// router.Use(muxlogrus.NewLogger().Middleware)

	// user endpoints
	user := router.PathPrefix("/user").Subrouter()
	user.HandleFunc("/add", controller.AddUserAPI).Methods("POST")
	user.HandleFunc("/get/{id}", controller.GetUserAPI).Methods("GET")
	user.HandleFunc("/delete/{id}", controller.DeleteUserAPI).Methods("DELETE")
	user.HandleFunc("/update", controller.UpdateUserAPI).Methods("PUT")

	// roles endpoints
	role := router.PathPrefix("/role").Subrouter()
	role.HandleFunc("/add", controller.AddRoleAPI).Methods("POST")
	role.HandleFunc("/get/{id}", controller.GetRoleAPI).Methods("GET")
	role.HandleFunc("/update", controller.UpdateRoleAPI).Methods("PUT")
	role.HandleFunc("/delete/{id}", controller.DeleteRoleAPI).Methods("DELETE")

	// user role endpoints
	userRole := user.PathPrefix("/role").Subrouter()
	userRole.HandleFunc("/associate", controller.AssociateUserRoleAPI).Methods("POST")
	userRole.HandleFunc("/disassociate", controller.DisassociateUserRoleAPI).Methods("DELETE")

	return router
}
