package db

import "project/model"

func (db *Connection) AddUserDB(user model.User) (model.User, error) {
	var resp model.User
	_, err := db.conn.NamedQuery(`INSERT INTO project.users(username,password,email,created_at)
	VALUES(:username,:password,:email,:created_at)`, user)
	if err != nil {
		return resp, err
	}
	return user, nil
}

func (db *Connection) AddRoleDB(role model.Role) (model.Role, error) {
	var resp model.Role
	_, err := db.conn.NamedQuery(`INSERT INTO project.roles(role_name)
	VALUES(:role_name)`, role)
	if err != nil {
		return resp, err
	}
	return role, nil
}

func (db *Connection) AssociateUserRoleDB(role model.RoleAssociate) (model.RoleAssociate, error) {
	var resp model.RoleAssociate
	_, err := db.conn.NamedQuery(`INSERT INTO project.user_roles(user_id,role_id)
	VALUES(:user_id,:role_id)`, role)
	if err != nil {
		return resp, err
	}
	return role, nil
}
