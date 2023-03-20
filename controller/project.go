package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"project/model"

	"github.com/gorilla/mux"
)

var AddProjectAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	var project model.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		log.Printf("decode error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	resp, err := conn.AddProjectDB(project)
	if err != nil {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var GetProjectAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	resp, err := conn.GetProjectByIDDB(ID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if resp.ID == "" {
		log.Printf("error: %v\n", "project doesn't exists")
		model.SuccessRespond("project doesn't exists", rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var GetAllProjectAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	resp, err := conn.GetAllProjectDB()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if len(resp) <= 0 {
		log.Printf("error: %v\n", "projects doesn't exists")
		model.SuccessRespond("projects doesn't exists", rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var UpdateProjectAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	var project model.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		log.Printf("decode error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	resp, err := conn.UpdateProjectDB(project)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("database error: %v\n", err.Error())
			model.ServerErrResponse("project doesn't exists", rw)
			return
		}
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var DeleteProjectAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	err := conn.DeleteProjectByIDDB(ID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond("project deleted successfully!", rw)
})
