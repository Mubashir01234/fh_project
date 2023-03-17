package model

type User struct {
	ID        string `json:"user_id,omitempty" db:"user_id"`
	Username  string `json:"username,omitempty" db:"username"`
	Password  string `json:"password,omitempty" db:"password"`
	Email     string `json:"email,omitempty" db:"email"`
	CreatedAt int64  `json:"created_at,omitempty" db:"created_at"`
}

type Role struct {
	ID       string `json:"role_id,omitempty" db:"role_id"`
	RoleName string `json:"role_name,omitempty" db:"role_name"`
}

type RoleAssociate struct {
	ID     string `json:"user_role_id,omitempty" db:"user_role_id"`
	UserID string `json:"user_id,omitempty" db:"user_id"`
	RoleID string `json:"role_id,omitempty" db:"role_id"`
}

type RoleDisassociate struct {
	ID     string `json:"user_role_id,omitempty" db:"user_role_id"`
	UserID string `json:"user_id,omitempty" db:"user_id"`
	RoleID string `json:"role_id,omitempty" db:"role_id"`
}
