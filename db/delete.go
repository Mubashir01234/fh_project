package db

func (db *Connection) DeleteUserByIDDB(ID string) error {
	_, err := db.conn.Exec(`DELETE FROM project.users WHERE user_id=?`, ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *Connection) DeleteRoleByIDDB(ID string) error {
	_, err := db.conn.Exec(`DELETE FROM project.roles WHERE role_id=?`, ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *Connection) DisassociateUserRoleDB(roleID, userID string) error {
	_, err := db.conn.Exec(`DELETE FROM project.user_roles WHERE role_id= ? AND user_id= ?`, roleID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (db *Connection) DeleteProjectByIDDB(ID string) error {
	_, err := db.conn.Exec(`DELETE FROM project.projects WHERE project_id=?`, ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *Connection) DeleteTaskByIDDB(ID string) error {
	_, err := db.conn.Exec(`DELETE FROM project.tasks WHERE task_id=?`, ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *Connection) DeleteTeamByIDDB(ID string) error {
	_, err := db.conn.Exec(`DELETE FROM project.teams WHERE team_id=?`, ID)
	if err != nil {
		return err
	}
	return nil
}
