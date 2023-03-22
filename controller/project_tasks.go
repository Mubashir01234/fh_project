package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"project/model"

	"github.com/gorilla/mux"
)

var AssignTaskToProjectAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	var body model.ProjectTask
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("decode error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	taskResp, err := conn.GetTaskByIDDB(body.TaskID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if len(taskResp.ID) <= 0 {
		log.Printf("error: %v\n", "task doesn't exists")
		model.ErrorResponse("task doesn't exist", rw)
		return
	}
	projectResp, err := conn.GetProjectByIDDB(body.ProjectID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if len(projectResp.ID) <= 0 {
		log.Printf("error: %v\n", "project doesn't exist")
		model.ServerErrResponse("project doesn't exist", rw)
		return
	}
	existedProjectID, err := conn.GetAssignTaskToProjectDB(body.ProjectID, body.TaskID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if existedProjectID == body.ProjectID {
		log.Printf("error: %v\n", "task already assign to this project")
		model.ErrorResponse("task already assign to this project", rw)
		return
	}
	resp, err := conn.AddProjectTaskDB(body)
	if err != nil {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var GetProjectTasksAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["id"]
	resp, err := conn.GetProjectTasksDB(projectID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if len(resp) <= 0 {
		log.Printf("error: %v\n", "project tasks doesn't exists")
		model.SuccessRespond("project tasks doesn't exists", rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var DeassignTaskFromProjectAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	projectTaskID := mux.Vars(r)["id"]
	err := conn.DeassignTaskFromProjectDB(projectTaskID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond("deassign task from project successfully", rw)
})
