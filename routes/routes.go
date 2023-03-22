package routes

import (
	"project/controller"

	_ "project/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
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
	user.HandleFunc("/", controller.GetAllUsersAPI).Methods("GET")

	// roles endpoints
	role := router.PathPrefix("/role").Subrouter()
	role.HandleFunc("/add", controller.AddRoleAPI).Methods("POST")
	role.HandleFunc("/get/{id}", controller.GetRoleAPI).Methods("GET")
	role.HandleFunc("/update", controller.UpdateRoleAPI).Methods("PUT")
	role.HandleFunc("/delete/{id}", controller.DeleteRoleAPI).Methods("DELETE")
	role.HandleFunc("/", controller.GetAllRolesAPI).Methods("GET")

	// user role endpoints
	userRole := user.PathPrefix("/role").Subrouter()
	userRole.HandleFunc("/associate", controller.AssociateUserRoleAPI).Methods("POST")
	userRole.HandleFunc("/disassociate", controller.DisassociateUserRoleAPI).Methods("DELETE")
	userRole.HandleFunc("/", controller.GetAllUsersRolesAPI).Methods("GET")

	// project endpoints
	project := router.PathPrefix("/project").Subrouter()
	project.HandleFunc("/add", controller.AddProjectAPI).Methods("POST")
	project.HandleFunc("/get/{id}", controller.GetProjectAPI).Methods("GET")
	project.HandleFunc("/", controller.GetAllProjectAPI).Methods("GET")
	project.HandleFunc("/update", controller.UpdateProjectAPI).Methods("PUT")
	project.HandleFunc("/delete/{id}", controller.DeleteProjectAPI).Methods("DELETE")

	// tasks endpoints
	task := router.PathPrefix("/task").Subrouter()
	task.HandleFunc("/add", controller.AddTaskAPI).Methods("POST")
	task.HandleFunc("/get/{id}", controller.GetTaskAPI).Methods("GET")
	task.HandleFunc("/", controller.GetAllTaskAPI).Methods("GET")
	task.HandleFunc("/update", controller.UpdateTaskAPI).Methods("PUT")
	task.HandleFunc("/delete/{id}", controller.DeleteTaskAPI).Methods("DELETE")

	// team endpoints
	team := router.PathPrefix("/team").Subrouter()
	team.HandleFunc("/add", controller.AddTeamAPI).Methods("POST")
	team.HandleFunc("/get/{id}", controller.GetTeamAPI).Methods("GET")
	team.HandleFunc("/", controller.GetAllTeamAPI).Methods("GET")
	team.HandleFunc("/update", controller.UpdateTeamAPI).Methods("PUT")
	team.HandleFunc("/delete/{id}", controller.DeleteTeamAPI).Methods("DELETE")

	// project tasks endpoint
	projectTask := project.PathPrefix("/task").Subrouter()
	projectTask.HandleFunc("/add", controller.AssignTaskToProjectAPI).Methods("POST")
	projectTask.HandleFunc("/get/{id}", controller.GetProjectTasksAPI).Methods("GET")
	projectTask.HandleFunc("/delete/{id}", controller.DeassignTaskFromProjectAPI).Methods("DELETE")

	// user team endpoints
	userTeam := team.PathPrefix("/user").Subrouter()
	userTeam.HandleFunc("/add", controller.AssignUserToTeamAPI).Methods("POST")
	userTeam.HandleFunc("/get/{id}", controller.GetUserTeamsAPI).Methods("GET")
	userTeam.HandleFunc("/delete/{id}", controller.DeassignUserFromTeamAPI).Methods("DELETE")

	// team project endpoints
	teamProject := project.PathPrefix("/team").Subrouter()
	teamProject.HandleFunc("/add", controller.AssignTeamToProjectAPI).Methods("POST")
	teamProject.HandleFunc("/get/{id}", controller.GetTeamProjectsAPI).Methods("GET")
	teamProject.HandleFunc("/delete/{id}", controller.DeassignTeamFromProjectAPI).Methods("DELETE")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}
