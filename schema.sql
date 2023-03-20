CREATE SCHEMA IF NOT EXISTS project;

CREATE TABLE IF NOT EXISTS project.users (
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS project.roles (
    role_id INT PRIMARY KEY AUTO_INCREMENT,
    role_name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS project.user_roles (
    user_role_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    role_id INT,
    CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES project.users(user_id),
    CONSTRAINT role_fk FOREIGN KEY (role_id) REFERENCES project.roles(role_id)
);

ALTER TABLE project.user_roles
	DROP CONSTRAINT user_fk;

ALTER TABLE project.user_roles    
    ADD CONSTRAINT user_fk
	FOREIGN KEY (user_id)
	REFERENCES project.users(user_id)
	ON DELETE CASCADE;

ALTER TABLE project.user_roles
	DROP CONSTRAINT role_fk;

ALTER TABLE project.user_roles    
    ADD CONSTRAINT role_fk
	FOREIGN KEY (role_id)
	REFERENCES project.roles(role_id)
	ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS project.projects (
    project_id INT PRIMARY KEY AUTO_INCREMENT,
    project_name VARCHAR(255) NOT NULL,
    project_description VARCHAR(255),
    start_date BIGINT NOT NULL,
    end_date BIGINT NOT NULL,
    budget VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS project.tasks (
    task_id INT PRIMARY KEY AUTO_INCREMENT,
    task_name VARCHAR(255) NOT NULL,
    task_description VARCHAR(255),
    start_date BIGINT NOT NULL,
    end_date BIGINT NOT NULL,
    status VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS project.project_tasks (
    project_task_id INT PRIMARY KEY AUTO_INCREMENT,
    project_id INT NOT NULL,
    task_id INT NOT NULL,
    CONSTRAINT project_fk FOREIGN KEY (project_id) REFERENCES project.projects(project_id) ON DELETE CASCADE,
    CONSTRAINT task_fk FOREIGN KEY (task_id) REFERENCES project.tasks(task_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS project.teams (
    team_id INT PRIMARY KEY AUTO_INCREMENT,
    team_name VARCHAR(255) NOT NULL,
    team_description VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS project.user_teams (
    user_teams_id INT PRIMARY KEY AUTO_INCREMENT,
    team_id INT NOT NULL,
    user_id INT NOT NULL,
    CONSTRAINT team_fk FOREIGN KEY (team_id) REFERENCES project.teams(team_id) ON DELETE CASCADE,
    CONSTRAINT user_id_fk FOREIGN KEY (user_id) REFERENCES project.users(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS project.team_projects (
    team_projects_id INT PRIMARY KEY AUTO_INCREMENT,
    team_id INT NOT NULL,
    project_id INT NOT NULL,
    CONSTRAINT project_team_fk FOREIGN KEY (team_id) REFERENCES project.teams(team_id) ON DELETE CASCADE,
    CONSTRAINT project_id_fk FOREIGN KEY (project_id) REFERENCES project.projects(project_id) ON DELETE CASCADE
);
