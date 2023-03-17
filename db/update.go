package db

import "project/model"

func (db *Connection) UpdateUserDB(user model.User) (model.User, error) {
	var resp model.User
	_, err := db.conn.NamedQuery(`UPDATE project.users SET username=:username,password=:password,email=:email
	WHERE user_id= :user_id`, user)
	if err != nil {
		return resp, err
	}
	err = db.conn.Get(&resp, `SELECT user_id,username,email FROM project.users WHERE user_id=?`, user.ID)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (db *Connection) UpdateRoleDB(role model.Role) (model.Role, error) {
	var resp model.Role
	_, err := db.conn.NamedQuery(`UPDATE project.roles SET role_name=:role_name
	WHERE role_id= :role_id`, role)
	if err != nil {
		return resp, err
	}
	err = db.conn.Get(&resp, `SELECT role_id,role_name FROM project.roles WHERE role_id=?`, role.ID)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
