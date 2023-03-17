package db

import "project/model"

func (db *Connection) GetUserDB(user model.User) (model.User, error) {
	var resp model.User
	err := db.conn.Get(&resp, `SELECT * FROM project.users WHERE email=?`, user.Email)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (db *Connection) GetUserByIDDB(ID string) (model.User, error) {
	var resp model.User
	err := db.conn.Get(&resp, `SELECT * FROM project.users WHERE user_id=?`, ID)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (db *Connection) GetRoleByRoleNameDB(role model.Role) (model.Role, error) {
	var resp model.Role
	err := db.conn.Get(&resp, `SELECT * FROM project.roles WHERE role_name=?`, role.RoleName)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (db *Connection) GetRoleByIDDB(ID string) (model.Role, error) {
	var resp model.Role
	err := db.conn.Get(&resp, `SELECT * FROM project.roles WHERE role_id=?`, ID)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (db *Connection) GetAssociateUserByUserIDDB(roleID, userID string) (string, error) {
	var ID string
	err := db.conn.Get(&ID, `SELECT role_id FROM project.user_roles WHERE role_id= ? AND user_id= ?`, roleID, userID)
	if err != nil {
		return "", err
	}
	return roleID, nil
}
