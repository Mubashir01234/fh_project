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

type Project struct {
	ID                 string `json:"project_id,omitempty" db:"project_id"`
	ProjectName        string `json:"project_name,omitempty" db:"project_name"`
	ProjectDescription string `json:"project_description,omitempty" db:"project_description"`
	StartDate          int64  `json:"start_date,omitempty" db:"start_date"`
	EndDate            int64  `json:"end_date,omitempty" db:"end_date"`
	Budget             string `json:"budget,omitempty" db:"budget"`
}

type Task struct {
	ID              string `json:"task_id,omitempty" db:"task_id"`
	TaskName        string `json:"task_name,omitempty" db:"task_name"`
	TaskDescription string `json:"task_description,omitempty" db:"task_description"`
	StartDate       int64  `json:"start_date,omitempty" db:"start_date"`
	EndDate         int64  `json:"end_date,omitempty" db:"end_date"`
	Status          string `json:"status,omitempty" db:"status"`
}

type Team struct {
	ID              string `json:"team_id,omitempty" db:"team_id"`
	TeamName        string `json:"team_name,omitempty" db:"team_name"`
	TeamDescription string `json:"team_description,omitempty" db:"team_description"`
}

type ProjectTask struct {
	ID        string `json:"project_task_id,omitempty" db:"project_task_id"`
	ProjectID string `json:"project_id,omitempty" db:"project_id"`
	TaskID    string `json:"task_id,omitempty" db:"task_id"`
}

type UserTeam struct {
	ID     string `json:"user_teams_id,omitempty" db:"user_teams_id"`
	TeamID string `json:"team_id,omitempty" db:"team_id"`
	UserID string `json:"user_id,omitempty" db:"user_id"`
}

type TeamProject struct {
	ID        string `json:"team_projects_id,omitempty" db:"team_projects_id"`
	TeamID    string `json:"team_id,omitempty" db:"team_id"`
	ProjectID string `json:"project_id,omitempty" db:"project_id"`
}
