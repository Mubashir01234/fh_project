package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"project/db"
	"project/model"
	"time"

	"github.com/gorilla/mux"
)

var conn = db.ConnectDB()

var AddUserAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("decode error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	user.CreatedAt = time.Now().Unix()
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		log.Printf("hash error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	existedUser, err := conn.GetUserDB(user)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if existedUser.Email != "" {
		log.Printf("error: %v\n", "email already exists")
		model.ErrorResponse("email already exists", rw)
		return
	}
	resp, err := conn.AddUserDB(user)
	if err != nil {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var GetUserAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	resp, err := conn.GetUserByIDDB(ID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if resp.ID == "" {
		log.Printf("error: %v\n", "user doesn't exists")
		model.SuccessRespond("user doesn't exists", rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var DeleteUserAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	err := conn.DeleteUserRoles(ID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	err = conn.DeleteUserByIDDB(ID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond("user deleted successfully!", rw)
})

var UpdateUserAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("decode error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	existedUser, err := conn.GetUserDB(user)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if existedUser.Email != "" {
		log.Printf("error: %v\n", "email already exists")
		model.ErrorResponse("email already exists", rw)
		return
	}
	if len(user.Password) > 0 {
		user.Password, err = HashPassword(user.Password)
		if err != nil {
			log.Printf("hash error: %v\n", err.Error())
			model.ServerErrResponse(err.Error(), rw)
			return
		}
	}
	resp, err := conn.UpdateUserDB(user)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("database error: %v\n", err.Error())
			model.ServerErrResponse("user doesn't exists", rw)
			return
		}
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond(resp, rw)
})
