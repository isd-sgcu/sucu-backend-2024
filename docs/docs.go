// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

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
        "/": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user profile",
                "parameters": [
                    {
                        "description": "Updated user data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.UpdateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/attachments": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attachments"
                ],
                "summary": "Get all attachments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.AttachmentDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/attachments/role/{role_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attachments"
                ],
                "summary": "Get all attachments by role",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Role of the user",
                        "name": "role_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.AttachmentDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/attachments/{attachment_id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attachments"
                ],
                "summary": "Delete an attachment by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Attachment ID",
                        "name": "attachment_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/attachments/{document_id}": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attachments"
                ],
                "summary": "Create new attachments",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Document ID",
                        "name": "document_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Attachment files",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Log in user",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "loginUserDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.LoginUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {},
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dtos.LoginResponseDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    }
                }
            }
        },
        "/auth/me": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Get current user profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {},
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dtos.UserDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/documents": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Documents"
                ],
                "summary": "Get all documents",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {},
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dtos.DocumentDTO"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Documents"
                ],
                "summary": "Create a new document",
                "parameters": [
                    {
                        "description": "Document data",
                        "name": "document",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateDocumentDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {}
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/documents/role/{role_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Documents"
                ],
                "summary": "Get documents by user role",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User role",
                        "name": "role_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {},
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dtos.DocumentDTO"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/documents/{document_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Documents"
                ],
                "summary": "Get document by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Document ID",
                        "name": "document_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {},
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dtos.DocumentDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Documents"
                ],
                "summary": "Update document by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Document ID",
                        "name": "document_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated document data",
                        "name": "document",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.UpdateDocumentDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Documents"
                ],
                "summary": "Delete document by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Document ID",
                        "name": "document_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/users": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {},
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dtos.UserDTO"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {}
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/users/{user_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {},
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dtos.UserDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated user data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.UpdateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.AttachmentDTO": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "document_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "role_id": {
                    "description": "role_id จะเอาไว้ให้ client ดูว่าไฟล์นี้มาจาก org อะไร sgcu or sucu",
                    "type": "string"
                },
                "type_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dtos.CreateDocumentDTO": {
            "type": "object",
            "properties": {
                "banner": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "cover": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "type_id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dtos.CreateUserDTO": {
            "type": "object",
            "required": [
                "first_name",
                "id",
                "last_name",
                "password"
            ],
            "properties": {
                "first_name": {
                    "description": "user's first name",
                    "type": "string"
                },
                "id": {
                    "description": "student id",
                    "type": "string"
                },
                "last_name": {
                    "description": "user's last name",
                    "type": "string"
                },
                "password": {
                    "description": "user's password",
                    "type": "string"
                },
                "role": {
                    "description": "role: sgcu-admin, sgcu-superadmin, sccu-admin, sccu-superadmin",
                    "type": "string"
                }
            }
        },
        "dtos.DocumentDTO": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/dtos.UserDTO"
                },
                "banner": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "cover": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "docs": {
                    "description": "docs file eg. pdf xlsx pptx",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.AttachmentDTO"
                    }
                },
                "id": {
                    "type": "string"
                },
                "images": {
                    "description": "images file eg. jpeg jpg png",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.AttachmentDTO"
                    }
                },
                "title": {
                    "type": "string"
                },
                "type_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dtos.LoginResponseDTO": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "dtos.LoginUserDTO": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "user's password",
                    "type": "string"
                },
                "student_id": {
                    "description": "user's id",
                    "type": "string"
                }
            }
        },
        "dtos.UpdateDocumentDTO": {
            "type": "object",
            "properties": {
                "banner": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "cover": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dtos.UpdateUserDTO": {
            "type": "object",
            "properties": {
                "first_name": {
                    "description": "user's first name",
                    "type": "string"
                },
                "last_name": {
                    "description": "user's last name",
                    "type": "string"
                },
                "password": {
                    "description": "user's password",
                    "type": "string"
                }
            }
        },
        "dtos.UserDTO": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "user's account creation time",
                    "type": "string"
                },
                "first_name": {
                    "description": "user's first name",
                    "type": "string"
                },
                "id": {
                    "description": "student id",
                    "type": "string"
                },
                "last_name": {
                    "description": "user's last name",
                    "type": "string"
                },
                "role": {
                    "description": "role: sgcu-admin, sgcu-superadmin, sccu-admin, sccu-superadmin",
                    "type": "string"
                },
                "updated_at": {
                    "description": "user's last update time",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type \"Bearer\" followed by a space and the token",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "SUCU Backend - API",
	Description:      "This is an SUCU Backend API in SUCU project.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
