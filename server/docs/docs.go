// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Batleforc",
            "url": "https://weebo.fr",
            "email": "maxleriche.60@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/asset/{fileName}": {
            "get": {
                "description": "Serve static asset",
                "tags": [
                    "Asset"
                ],
                "summary": "Serve static asset",
                "parameters": [
                    {
                        "type": "string",
                        "description": "fileName",
                        "name": "fileName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Login body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/route.LoginBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login return",
                        "schema": {
                            "$ref": "#/definitions/route.LoginReturn"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "description": "Logout user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout user",
                "parameters": [
                    {
                        "description": "Logout body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/route.LogoutBody"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register User, Email has to be Unique and valid, Pseudo has to be Unique and \u003e 3 characters, Password has to be \u003e 8 characters, Name and surname has to be \u003e 2 characters",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register User",
                "parameters": [
                    {
                        "description": "Register body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/route.RegisterBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Register return",
                        "schema": {
                            "$ref": "#/definitions/route.RegisterReturn"
                        }
                    }
                }
            }
        },
        "/auth/renew": {
            "post": {
                "description": "Renew Token via refresh token",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Renew Token",
                "parameters": [
                    {
                        "description": "Renew body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/route.RenewTokenBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Renew return",
                        "schema": {
                            "$ref": "#/definitions/route.RenewTokenReturn"
                        }
                    }
                }
            }
        },
        "/chan": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create channel, Name has to be unique",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Chan"
                ],
                "summary": "Create channel",
                "parameters": [
                    {
                        "description": "Create channel body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/route.CreateChanBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Create channel return",
                        "schema": {
                            "$ref": "#/definitions/route.CreateChanReturn"
                        }
                    },
                    "400": {
                        "description": "Create channel return",
                        "schema": {
                            "$ref": "#/definitions/route.CreateChanReturn"
                        }
                    },
                    "500": {
                        "description": "Create channel return",
                        "schema": {
                            "$ref": "#/definitions/route.CreateChanReturn"
                        }
                    }
                }
            }
        },
        "/chan/name": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Check if channel name is available",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Chan"
                ],
                "summary": "Check if channel name is available",
                "parameters": [
                    {
                        "description": "Check channel name body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/route.CheckChanNameBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Chann available or not",
                        "schema": {
                            "$ref": "#/definitions/route.CheckChanNameReturn"
                        }
                    },
                    "400": {
                        "description": "Body not valid",
                        "schema": {
                            "$ref": "#/definitions/route.CheckChanNameReturn"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/route.CheckChanNameReturn"
                        }
                    }
                }
            }
        },
        "/chan/{chanId}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get One Channel by id",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Chan"
                ],
                "summary": "Get One Channel by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Channel id",
                        "name": "chanId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Channel",
                        "schema": {
                            "$ref": "#/definitions/model.Channel"
                        }
                    },
                    "400": {
                        "description": "Chan Id is not valid",
                        "schema": {
                            "$ref": "#/definitions/route.GetOneChanReturn"
                        }
                    },
                    "403": {
                        "description": "User is not in channel",
                        "schema": {
                            "$ref": "#/definitions/route.GetOneChanReturn"
                        }
                    },
                    "500": {
                        "description": "Error while getting channel, (can be normal if not exist)",
                        "schema": {
                            "$ref": "#/definitions/route.GetOneChanReturn"
                        }
                    }
                }
            }
        },
        "/chan/{chanId}/message": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get One Channel message by id, if user not in chan can't see message",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Chan"
                ],
                "summary": "Get One Channel message by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Channel id",
                        "name": "chanId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Message"
                            }
                        }
                    },
                    "400": {
                        "description": "Chan Id is not valid",
                        "schema": {
                            "$ref": "#/definitions/route.GetOneChanReturn"
                        }
                    },
                    "403": {
                        "description": "User is not in channel",
                        "schema": {
                            "$ref": "#/definitions/route.GetOneChanReturn"
                        }
                    },
                    "500": {
                        "description": "Error while getting channel, (can be normal if not exist)",
                        "schema": {
                            "$ref": "#/definitions/route.GetOneChanReturn"
                        }
                    }
                }
            }
        },
        "/chan/{chanId}/renew": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Reset channel password",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Chan"
                ],
                "summary": "Reset channel password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Channel id",
                        "name": "chanId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Reset channel password return",
                        "schema": {
                            "$ref": "#/definitions/route.RenewChanPasswordReturn"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get user",
                "tags": [
                    "User"
                ],
                "summary": "Get user",
                "responses": {
                    "200": {
                        "description": "user return",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Set user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Set user",
                "parameters": [
                    {
                        "description": "Set user body",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/route.SetUserBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Set user return",
                        "schema": {
                            "$ref": "#/definitions/route.SetUserReturn"
                        }
                    }
                }
            }
        },
        "/user/setpicture": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "SetPicture user",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "User"
                ],
                "summary": "SetPicture user",
                "parameters": [
                    {
                        "type": "file",
                        "description": ".jpeg, .png, .gif",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "SetPicture return",
                        "schema": {
                            "$ref": "#/definitions/route.SetPictureReturn"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "model.Channel": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "type": "integer"
                },
                "picture": {
                    "type": "string"
                },
                "private": {
                    "type": "boolean"
                },
                "updatedAt": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ChannelUser"
                    }
                }
            }
        },
        "model.ChannelUser": {
            "type": "object",
            "properties": {
                "canMod": {
                    "type": "boolean"
                },
                "canRead": {
                    "type": "boolean"
                },
                "canSend": {
                    "type": "boolean"
                },
                "channelID": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.PublicUser"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "model.Message": {
            "type": "object",
            "properties": {
                "channelID": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "model.PublicUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "pseudo": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "channels": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ChannelUser"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "myChannels": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Channel"
                    }
                },
                "name": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "pseudo": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "route.CheckChanNameBody": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "route.CheckChanNameReturn": {
            "type": "object",
            "properties": {
                "available": {
                    "type": "boolean"
                },
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "route.CreateChanBody": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "private": {
                    "type": "boolean"
                }
            }
        },
        "route.CreateChanReturn": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "passkey": {
                    "type": "string"
                },
                "updated": {
                    "type": "boolean"
                }
            }
        },
        "route.GetOneChanReturn": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "route.LoginBody": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "route.LoginReturn": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "pseudo": {
                    "type": "string"
                },
                "renew_token": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "route.LogoutBody": {
            "type": "object",
            "properties": {
                "renew_token": {
                    "type": "string"
                }
            }
        },
        "route.RegisterBody": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "pseudo": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "route.RegisterReturn": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "registered": {
                    "type": "boolean"
                }
            }
        },
        "route.RenewChanPasswordReturn": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "updated": {
                    "type": "boolean"
                }
            }
        },
        "route.RenewTokenBody": {
            "type": "object",
            "properties": {
                "renew_token": {
                    "type": "string"
                }
            }
        },
        "route.RenewTokenReturn": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "pseudo": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "route.SetPictureReturn": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "route.SetUserBody": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "route.SetUserReturn": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "updated": {
                    "type": "boolean"
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
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Bipper Api",
	Description:      "Bipper api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
