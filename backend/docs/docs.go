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
        "/auth/google": {
            "post": {
                "description": "Authenticate user with Google OAuth token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Google Auth",
                "parameters": [
                    {
                        "description": "Google OAuth token",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.googleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User info from Google",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request - invalid or missing Google OAuth token",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error - failed to process the request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway - incorrect login method for the user's account provider",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/signin": {
            "post": {
                "description": "Login user with the provided credentials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.loginUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User login successful",
                        "schema": {
                            "$ref": "#/definitions/api.loginUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request - invalid input details or wrong password",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not found - user with the given email does not exist",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway - incorrect login method for the user's account provider",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Register a new user with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.registerUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User registration successful",
                        "schema": {
                            "$ref": "#/definitions/api.registerUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request - invalid input details",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden - email is already registered",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/chats/histories": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves the chat histories for the logged-in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Get chat histories",
                "responses": {
                    "200": {
                        "description": "Successful retrieval of chat histories",
                        "schema": {
                            "$ref": "#/definitions/api.getChatsHistoriesResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Invalid or missing JWT token",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/chats/{uid}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves the details of the chat history between the logged-in user and the user specified by UID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Get chat details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID to get chat details with",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful retrieval of chat details",
                        "schema": {
                            "$ref": "#/definitions/api.getChatsDetailsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request - Invalid UID",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Invalid or missing JWT token",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found - User not found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    },
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieves a list of social users excluding the current user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get social users",
                "responses": {
                    "200": {
                        "description": "Successful retrieval of users data",
                        "schema": {
                            "$ref": "#/definitions/api.getSocialUserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized access",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Chats": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "receiver": {
                    "$ref": "#/definitions/api.UserResponse"
                },
                "receiver_uid": {
                    "type": "string"
                },
                "sender": {
                    "$ref": "#/definitions/api.UserResponse"
                },
                "sender_uid": {
                    "type": "string"
                }
            }
        },
        "api.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "api.History": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "latest_content": {
                    "type": "string"
                },
                "latest_created_at": {
                    "type": "string"
                },
                "receiver_name": {
                    "type": "string"
                },
                "receiver_uid": {
                    "type": "string"
                },
                "sender_name": {
                    "type": "string"
                },
                "sender_uid": {
                    "type": "string"
                }
            }
        },
        "api.UserResponse": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "provider": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                }
            }
        },
        "api.getChatsDetailsResponse": {
            "type": "object",
            "properties": {
                "chats": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Chats"
                    }
                },
                "message": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/api.UserResponse"
                }
            }
        },
        "api.getChatsHistoriesResponse": {
            "type": "object",
            "properties": {
                "histories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.History"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "api.getSocialUserResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.UserResponse"
                    }
                }
            }
        },
        "api.googleRequest": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "api.loginUserRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "api.loginUserResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/api.UserResponse"
                }
            }
        },
        "api.registerUserRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "api.registerUserResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/api.UserResponse"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Chitchat API Documentation",
	Description:      "This is a documentation for Chitchat API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
