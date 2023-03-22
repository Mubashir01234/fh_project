package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"project/model"

	"github.com/gorilla/mux"
)

var AssignTeamToProjectAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	var body model.TeamProject
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("decode error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
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
		model.ErrorResponse("project doesn't exist", rw)
		return
	}
	teamResp, err := conn.GetTeamByIDDB(body.TeamID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if len(teamResp.ID) <= 0 {
		log.Printf("error: %v\n", "team doesn't exist")
		model.ServerErrResponse("team doesn't exist", rw)
		return
	}
	existedTeamID, err := conn.GetAssignTeamToProjectDB(body.TeamID, body.ProjectID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if existedTeamID == body.TeamID {
		log.Printf("error: %v\n", "team already assign to this project")
		model.ErrorResponse("team already assign to this project", rw)
		return
	}
	resp, err := conn.AddTeamToProjectDB(body)
	if err != nil {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var GetTeamProjectsAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	teamID := mux.Vars(r)["id"]
	resp, err := conn.GetTeamProjectsDB(teamID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if len(resp) <= 0 {
		log.Printf("error: %v\n", "team projects doesn't exist")
		model.SuccessRespond("team projects doesn't exist", rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var DeassignTeamFromProjectAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	projectTeamID := mux.Vars(r)["id"]
	err := conn.DeassignTeamFromProjectDB(projectTeamID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond("deassign team from project successfully", rw)
})
