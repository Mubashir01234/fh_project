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
