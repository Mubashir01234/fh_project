package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"project/model"

	"github.com/gorilla/mux"
)

var AddRoleAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	var role model.Role
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		log.Printf("decode error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	existedRole, err := conn.GetRoleByRoleNameDB(role)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if existedRole.RoleName == role.RoleName {
		log.Printf("error: %v\n", "role already exists")
		model.ErrorResponse("role already exists", rw)
		return
	}
	resp, err := conn.AddRoleDB(role)
	if err != nil {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var GetRoleAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	resp, err := conn.GetRoleByIDDB(ID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if resp.ID == "" {
		log.Printf("error: %v\n", "role doesn't exists")
		model.SuccessRespond("role doesn't exists", rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var UpdateRoleAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	var role model.Role
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		log.Printf("decode error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	existedRole, err := conn.GetRoleByRoleNameDB(role)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if existedRole.RoleName == role.RoleName {
		log.Printf("error: %v\n", "role already exists")
		model.ErrorResponse("role already exists", rw)
		return
	}
	resp, err := conn.UpdateRoleDB(role)
	if err != nil {
		log.Printf("database error: %v\n", err.Error())
		if err == sql.ErrNoRows {
			log.Printf("error: %v\n", "role already exists")
			model.ServerErrResponse("role doesn't exists", rw)
			return
		}
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var DeleteRoleAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	err := conn.DeleteRoleByIDDB(ID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond("role deleted successfully!", rw)
})

var AssociateUserRoleAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	var roleAssociate model.RoleAssociate
	err := json.NewDecoder(r.Body).Decode(&roleAssociate)
	if err != nil {
		log.Printf("decode error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	existedRoleID, err := conn.GetAssociateUserByUserIDDB(roleAssociate.RoleID, roleAssociate.UserID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if existedRoleID == roleAssociate.RoleID {
		log.Printf(" error: %v\n", "user already associate with this role")
		model.ErrorResponse("user already associate with this role", rw)
		return
	}
	userResp, err := conn.GetUserByIDDB(roleAssociate.UserID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if userResp.ID != roleAssociate.UserID {
		log.Printf("error: %v\n", "user doesn't exists")
		model.ServerErrResponse("user doesn't exists", rw)
		return
	}
	roleResp, err := conn.GetRoleByIDDB(roleAssociate.RoleID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if roleResp.ID != roleAssociate.RoleID {
		log.Printf("error: %v\n", "role doesn't exists")
		model.ServerErrResponse("role doesn't exists", rw)
		return
	}
	resp, err := conn.AssociateUserRoleDB(roleAssociate)
	if err != nil {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var DisassociateUserRoleAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	var roleDisAssociate model.RoleDisassociate
	err := json.NewDecoder(r.Body).Decode(&roleDisAssociate)
	if err != nil {
		log.Printf("decode error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}

	existedRoleID, err := conn.GetAssociateUserByUserIDDB(roleDisAssociate.RoleID, roleDisAssociate.UserID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if existedRoleID == "" {
		log.Printf("error: %v\n", "user doesn't associated with this role")
		model.ErrorResponse("user doesn't associated with this role", rw)
		return
	}
	err = conn.DisassociateUserRoleDB(roleDisAssociate.RoleID, roleDisAssociate.UserID)
	if err != nil {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	model.SuccessRespond("disaccociate role successfully", rw)
})

var GetAllRolesAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	resp, err := conn.GetAllRolesDB()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if len(resp) <= 0 {
		log.Printf("error: %v\n", "roles doesn't exists")
		model.SuccessRespond("roles doesn't exists", rw)
		return
	}
	model.SuccessRespond(resp, rw)
})

var GetAllUsersRolesAPI = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	resp, err := conn.GetAllUsersRolesDB()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("database error: %v\n", err.Error())
		model.ServerErrResponse(err.Error(), rw)
		return
	}
	if len(resp) <= 0 {
		log.Printf("error: %v\n", "users roles doesn't exists")
		model.SuccessRespond("users roles doesn't exists", rw)
		return
	}
	model.SuccessRespond(resp, rw)
})
