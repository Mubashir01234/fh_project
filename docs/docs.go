// Code generated by swaggo/swag. DO NOT EDIT
package docs

import (
    "github.com/swaggo/swag"
    "project/handlers"
    
)
    

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/add": {
            "post": {
                summary: Add new user,
                tags:[
                    User
                ],
                produces: [
					application/json
				],
                parameters: [
                    {
                        name: add user,
                        in: body,   
                        "properties":{
                            email:{
                                type:string,
                                format:json
                            },
                            username:{
                                type:string,
                                format:json
                            },
                            password:{
                                type:string,
                                format:json
                            },
                        },
			        },
			    ],
                responses: {
					200: {
						description: Status, OK
					}
				}
            }
        },
        "/user/get/{id}":{
            "get":{
                summary: Get user by ID,
                tags:[
                    User
                ],
                produces: [
					application/json
				],
                parameters: [
					{
						name: id,
						in: path,
						description: give user id,
						required: true,
						type: string
					},
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/user/":{
            "get":{
                summary: Get all users,
                tags:[
                    User
                ],
                produces: [
					application/json
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/user/update":{
            "put":{
                summary: Update user,
                tags:[
                    User
                ],
                produces: [
					application/json
				],
                parameters: [
                    {
                        name: update user,
                        in: body,
                        "properties":{
                            user_id:{
                                type:string,
                                format:json
                            },
                            email:{
                                type:string,
                                format:json
                            },
                            username:{
                                type:string,
                                format:json
                            },
                            password:{
                                type:string,
                                format:json
                            },
                        },
                    }
			    ],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/user/delete/{id}":{
            "delete":{
                summary: Delete user by ID,
                tags:[
                    User
                ],
                produces: [
					application/json
				],
                parameters: [
					{
						name: id,
						in: path,
						description: give user id,
						required: true,
						type: string
					},
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/role/add": {
            "post": {
                summary: Add new role,
                tags:[
                    Role
                ],
                produces: [
					application/json
				],
                parameters: [
                    {
                        name: add role,
                        in: body,   
                        "properties":{
                            role_name:{
                                type:string,
                                format:json
                            }
                        },
			        },
			    ],
                responses: {
					200: {
						description: Status, OK
					}
				}
            }
        },
        "/role/get/{id}":{
            "get":{
                summary: Get role by ID,
                tags:[
                    Role
                ],
                produces: [
					application/json
				],
                parameters: [
					{
						name: id,
						in: path,
						description: give user id,
						required: true,
						type: string
					},
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/role/":{
            "get":{
                summary: Get all roles,
                tags:[
                    Role
                ],
                produces: [
					application/json
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/role/update":{
            "put":{
                summary: Update role,
                tags:[
                    Role
                ],
                produces: [
					application/json
				],
                parameters: [
                    {
                        name: update role,
                        in: body,
                        "properties":{
                            role_id:{
                                type:string,
                                format:json
                            },
                            role_name:{
                                type:string,
                                format:json
                            }
                        },
                    }
			    ],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/role/delete/{id}":{
            "delete":{
                summary: Delete role by ID,
                tags:[
                    Role
                ],
                produces: [
					application/json
				],
                parameters: [
					{
						name: id,
						in: path,
						description: give role id,
						required: true,
						type: string
					},
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/user/role/associate": {
            "post": {
                summary: User to role associate,
                tags:[
                    Role associate
                ],
                produces: [
					application/json
				],
                parameters: [
                    {
                        name: associate role,
                        in: body,   
                        "properties":{
                            user_id:{
                                type:string,
                                format:json
                            },
                            role_id:{
                                type:string,
                                format:json
                            }
                        },
			        },
			    ],
                responses: {
					200: {
						description: Status, OK
					}
				}
            }
        },
        "/user/role/":{
            "get":{
                summary: Get all users roles,
                tags:[
                    Role associate
                ],
                produces: [
					application/json
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/user/role/disassociate":{
            "delete":{
                summary: User to role disassociate,
                tags:[
                    Role associate
                ],
                produces: [
					application/json
				],
                parameters: [
                    {
                        name: disassociate role,
                        in: body,   
                        "properties":{
                            user_id:{
                                type:string,
                                format:json
                            },
                            role_id:{
                                type:string,
                                format:json
                            }
                        },
			        },
			    ],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/project/add": {
            "post": {
                summary: Add new project,
                tags:[
                    Project
                ],
                produces: [
					application/json
				],
                parameters: [
                    {
                        name: add project,
                        in: body,   
                        "properties":{
                            project_name:{
                                type:string,
                                format:json
                            },
                            project_description:{
                                type:string,
                                format:json
                            },
                            start_date:{
                                type:int64,
                                format:json
                            },
                            end_date:{
                                type:int64,
                                format:json
                            },
                            budget:{
                                type:string,
                                format:json
                            }
                        },
			        },
			    ],
                responses: {
					200: {
						description: Status, OK
					}
				}
            }
        },
        "/project/get/{id}":{
            "get":{
                summary: Get project by ID,
                tags:[
                    Project
                ],
                produces: [
					application/json
				],
                parameters: [
					{
						name: id,
						in: path,
						description: give project id,
						required: true,
						type: string
					},
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/project/":{
            "get":{
                summary: Get all projects,
                tags:[
                    Project
                ],
                produces: [
					application/json
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/project/update":{
            "put":{
                summary: Update project,
                tags:[
                    Project
                ],
                produces: [
					application/json
				],
                parameters: [
                    {
                        name: update project,
                        in: body,
                        "properties":{
                            project_id:{
                                type:string,
                                format:json
                            },
                            project_name:{
                                type:string,
                                format:json
                            },
                            project_description:{
                                type:string,
                                format:json
                            },
                            start_date:{
                                type:int64,
                                format:json
                            },
                            end_date:{
                                type:int64,
                                format:json
                            },
                            budget:{
                                type:string,
                                format:json
                            }
                        },
                    }
			    ],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/project/delete/{id}":{
            "delete":{
                summary: Delete project by ID,
                tags:[
                    Project
                ],
                produces: [
					application/json
				],
                parameters: [
					{
						name: id,
						in: path,
						description: give project id,
						required: true,
						type: string
					},
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/task/add": {
            "post": {
                summary: Add new task,
                tags:[
                    Task
                ],
                produces: [
					application/json
				],
                parameters: [
                    {
                        name: add project,
                        in: body,   
                        "properties":{
                            task_name:{
                                type:string,
                                format:json
                            },
                            task_description:{
                                type:string,
                                format:json
                            },
                            start_date:{
                                type:int64,
                                format:json
                            },
                            end_date:{
                                type:int64,
                                format:json
                            },
                            status:{
                                type:string,
                                format:json
                            }
                        },
			        },
			    ],
                responses: {
					200: {
						description: Status, OK
					}
				}
            }
        },
        "/task/get/{id}":{
            "get":{
                summary: Get task by ID,
                tags:[
                    Task
                ],
                produces: [
					application/json
				],
                parameters: [
					{
						name: id,
						in: path,
						description: give task id,
						required: true,
						type: string
					},
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/task/":{
            "get":{
                summary: Get all tasks,
                tags:[
                    Task
                ],
                produces: [
					application/json
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/task/update":{
            "put":{
                summary: Update task,
                tags:[
                    Task
                ],
                produces: [
					application/json
				],
                parameters: [
                    {
                        name: update project,
                        in: body,
                        "properties":{
                            task_id:{
                                type:string,
                                format:json
                            },
                            task_name:{
                                type:string,
                                format:json
                            },
                            task_description:{
                                type:string,
                                format:json
                            },
                            start_date:{
                                type:int64,
                                format:json
                            },
                            end_date:{
                                type:int64,
                                format:json
                            },
                            status:{
                                type:string,
                                format:json
                            }
                        },
                    }
			    ],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/task/delete/{id}":{
            "delete":{
                summary: Delete task by ID,
                tags:[
                    Task
                ],
                produces: [
					application/json
				],
                parameters: [
					{
						name: id,
						in: path,
						description: give project id,
						required: true,
						type: string
					},
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/team/add": {
            "post": {
                summary: Add new team,
                tags:[
                    Team
                ],
                produces: [
					application/json
				],
                parameters: [
                    {
                        name: add team,
                        in: body,   
                        "properties":{
                            team_name:{
                                type:string,
                                format:json
                            },
                            team_description:{
                                type:string,
                                format:json
                            }
                        },
			        },
			    ],
                responses: {
					200: {
						description: Status, OK
					}
				}
            }
        },
        "/team/get/{id}":{
            "get":{
                summary: Get team by ID,
                tags:[
                    Team
                ],
                produces: [
					application/json
				],
                parameters: [
					{
						name: id,
						in: path,
						description: give team id,
						required: true,
						type: string
					},
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/team/":{
            "get":{
                summary: Get all teams,
                tags:[
                    Team
                ],
                produces: [
					application/json
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/team/update":{
            "put":{
                summary: Update team,
                tags:[
                    Team
                ],
                produces: [
					application/json
				],
                parameters: [
                    {
                        name: update team,
                        in: body,
                        "properties":{
                            team_id:{
                                type:string,
                                format:json
                            },
                            team_name:{
                                type:string,
                                format:json
                            },
                            team_description:{
                                type:string,
                                format:json
                            }
                        },
                    }
			    ],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/team/delete/{id}":{
            "delete":{
                summary: Delete team by ID,
                tags:[
                    Team
                ],
                produces: [
					application/json
				],
                parameters: [
					{
						name: id,
						in: path,
						description: give team id,
						required: true,
						type: string
					},
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/project/task/add": {
            "post": {
                summary: Assign task to project,
                tags:[
                    Project Task
                ],
                produces: [
					application/json
				],
                parameters: [
                    {
                        name: assign task to project,
                        in: body,   
                        "properties":{
                            project_id:{
                                type:string,
                                format:json
                            },
                            task_id:{
                                type:string,
                                format:json
                            }
                        },
			        },
			    ],
                responses: {
					200: {
						description: Status, OK
					}
				}
            }
        },
        "/project/task/get/{id}":{
            "get":{
                summary: Get project tasks by project ID,
                tags:[
                    Project Task
                ],
                produces: [
					application/json
				],
                parameters: [
					{
						name: id,
						in: path,
						description: give project_id,
						required: true,
						type: string
					},
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/project/task/delete/{id}":{
            "delete":{
                summary: Deassign task from project by project task id,
                tags:[
                    Project Task
                ],
                produces: [
					application/json
				],
                parameters: [
					{
						name: id,
						in: path,
						description: give project task id,
						required: true,
						type: string
					},
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/team/user/add": {
            "post": {
                summary: Assign user to team,
                tags:[
                    User Team
                ],
                produces: [
					application/json
				],
                parameters: [
                    {
                        name: assign user to team,
                        in: body,   
                        "properties":{
                            team_id:{
                                type:string,
                                format:json
                            },
                            user_id:{
                                type:string,
                                format:json
                            }
                        },
			        },
			    ],
                responses: {
					200: {
						description: Status, OK
					}
				}
            }
        },
        "/team/user/get/{id}":{
            "get":{
                summary: Get team user by team ID,
                tags:[
                    User Team
                ],
                produces: [
					application/json
				],
                parameters: [
					{
						name: id,
						in: path,
						description: give team_id,
						required: true,
						type: string
					},
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
        "/team/user/delete/{id}":{
            "delete":{
                summary: Deassign user from team by user team id",
                tags:[
                    User Team
                ],
                produces: [
					application/json
				],
                parameters: [
					{
						name: id,
						in: path,
						description: give project task id,
						required: true,
						type: string
					},
				],
                responses: {
					200: {
						description: Status, OK
					}
				}                
            }
        },
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:"+handlers.DotEnvVariable("SERVER_PORT"),
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "FH_Project",
	Description:      "Swagger for FH_Project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
