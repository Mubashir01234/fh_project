package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"project/model"

	"github.com/gorilla/mux"
)

var AddTeamAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	var team model.Team
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		log.Printf("decode error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	resp, err := conn.AddTeamDB(team)
	if err != nil {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var GetTeamAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	resp, err := conn.GetTeamByIDDB(ID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if resp.ID == "" {
		log.Printf("error: %v\n", "team doesn't exists")
		model.SuccessRespond("team doesn't exists", rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var GetAllTeamAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	resp, err := conn.GetAllTeamDB()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if len(resp) <= 0 {
		log.Printf("error: %v\n", "teams doesn't exists")
		model.SuccessRespond("teams doesn't exists", rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var UpdateTeamAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	var team model.Team
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		log.Printf("decode error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	resp, err := conn.UpdateTeamDB(team)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("database error: %v\n", err.Error())
			model.ServerErrResponse("team doesn't exists", rw)
			return
		}
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var DeleteTeamAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	err := conn.DeleteTeamByIDDB(ID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond("team deleted successfully!", rw)
})
