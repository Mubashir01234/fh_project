package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"project/model"

	"github.com/gorilla/mux"
)

var AddTaskAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Printf("decode error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	resp, err := conn.AddTaskDB(task)
	if err != nil {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var GetTaskAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	resp, err := conn.GetTaskByIDDB(ID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if resp.ID == "" {
		log.Printf("error: %v\n", "task doesn't exists")
		model.SuccessRespond("task doesn't exists", rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var GetAllTaskAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	resp, err := conn.GetAllTaskDB()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if len(resp) <= 0 {
		log.Printf("error: %v\n", "tasks doesn't exists")
		model.SuccessRespond("tasks doesn't exists", rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var UpdateTaskAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Printf("decode error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	resp, err := conn.UpdateTaskDB(task)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("database error: %v\n", err.Error())
			model.ServerErrResponse("task doesn't exists", rw)
			return
		}
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var DeleteTaskAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	err := conn.DeleteTaskByIDDB(ID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond("task deleted successfully!", rw)
})
