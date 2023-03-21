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

func (db *Connection) GetAllUsersDB() ([]model.User, error) {
	var resp []model.User
	err := db.conn.Select(&resp, `SELECT * FROM project.users`)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (db *Connection) GetAllRolesDB() ([]model.Role, error) {
	var resp []model.Role
	err := db.conn.Select(&resp, `SELECT * FROM project.roles`)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (db *Connection) GetAllUsersRolesDB() ([]model.RoleAssociate, error) {
	var resp []model.RoleAssociate
	err := db.conn.Select(&resp, `SELECT * FROM project.user_roles`)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (db *Connection) GetProjectByIDDB(ID string) (model.Project, error) {
	var resp model.Project
	err := db.conn.Get(&resp, `SELECT * FROM project.projects WHERE project_id=?`, ID)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (db *Connection) GetAllProjectDB() ([]model.Project, error) {
	var resp []model.Project
	err := db.conn.Select(&resp, `SELECT * FROM project.projects`)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (db *Connection) GetTaskByIDDB(ID string) (model.Task, error) {
	var resp model.Task
	err := db.conn.Get(&resp, `SELECT * FROM project.tasks WHERE task_id=?`, ID)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (db *Connection) GetAllTaskDB() ([]model.Task, error) {
	var resp []model.Task
	err := db.conn.Select(&resp, `SELECT * FROM project.tasks`)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (db *Connection) GetTeamByIDDB(ID string) (model.Team, error) {
	var resp model.Team
	err := db.conn.Get(&resp, `SELECT * FROM project.teams WHERE team_id=?`, ID)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (db *Connection) GetAllTeamDB() ([]model.Team, error) {
	var resp []model.Team
	err := db.conn.Select(&resp, `SELECT * FROM project.teams`)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
