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

