package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"project/model"

	"github.com/gorilla/mux"
)

var AssignUserToTeamAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	var body model.UserTeam
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("decode error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	userResp, err := conn.GetUserByIDDB(body.UserID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if len(userResp.ID) <= 0 {
		log.Printf("error: %v\n", "user doesn't exist")
		model.ErrorResponse("user doesn't exist", rw)
		return
	}
	teamResp, err := conn.GetTeamByIDDB(body.TeamID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if len(teamResp.ID) <= 0 {
		log.Printf("error: %v\n", "team doesn't exists")
		model.ServerErrResponse("team doesn't exists", rw)
		return
	}
	resp, err := conn.AddUserToTeamDB(body)
	if err != nil {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var GetUserTeamsAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	teamID := mux.Vars(r)["id"]
	resp, err := conn.GetTeamUsersDB(teamID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if len(resp) <= 0 {
		log.Printf("error: %v\n", "team users doesn't exist")
		model.SuccessRespond("team users doesn't exist", rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var DeassignUserFromTeamAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	userTeamID := mux.Vars(r)["id"]
	err := conn.DeassignUserFromTeamDB(userTeamID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond("deassign task from project successfully", rw)
})
