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

func (db *Connection) AddProjectDB(project model.Project) (model.Project, error) {
	var resp model.Project
	_, err := db.conn.NamedQuery(`INSERT INTO project.projects(project_name,project_description,start_date,end_date,budget)
	VALUES(:project_name,:project_description,:start_date,:end_date,:budget)`, project)
	if err != nil {
		return resp, err
	}
	return project, nil
}

func (db *Connection) AddTaskDB(task model.Task) (model.Task, error) {
	var resp model.Task
	_, err := db.conn.NamedQuery(`INSERT INTO project.tasks(task_name,task_description,start_date,end_date,status)
	VALUES(:task_name,:task_description,:start_date,:end_date,:status)`, task)
	if err != nil {
		return resp, err
	}
	return task, nil
}

func (db *Connection) AddTeamDB(team model.Team) (model.Team, error) {
	var resp model.Team
	_, err := db.conn.NamedQuery(`INSERT INTO project.teams(team_name,team_description)
	VALUES(:team_name,:team_description)`, team)
	if err != nil {
		return resp, err
	}
	return team, nil
}

func (db *Connection) AddProjectTaskDB(project model.ProjectTask) (model.ProjectTask, error) {
	var resp model.ProjectTask
	_, err := db.conn.NamedQuery(`INSERT INTO project.project_tasks(project_id,task_id)
	VALUES(:project_id,:task_id)`, project)
	if err != nil {
		return resp, err
	}
	return project, nil
}

func (db *Connection) AddUserToTeamDB(user model.UserTeam) (model.UserTeam, error) {
	var resp model.UserTeam
	_, err := db.conn.NamedQuery(`INSERT INTO project.user_teams(team_id,user_id)
	VALUES(:team_id,:user_id)`, user)
	if err != nil {
		return resp, err
	}
	return user, nil
}

func (db *Connection) AddTeamToProjectDB(data model.TeamProject) (model.TeamProject, error) {
	var resp model.TeamProject
	_, err := db.conn.NamedQuery(`INSERT INTO project.team_projects(team_id,project_id)
	VALUES(:team_id,:project_id)`, data)
	if err != nil {
		return resp, err
	}
	return data, nil
}
