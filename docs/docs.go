// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/announcements/register": {
            "post": {
                "description": "Register a new announcement",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "announcements"
                ],
                "summary": "Register a new announcement",
                "parameters": [
                    {
                        "description": "Announcement",
                        "name": "announcement",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Announcement"
                        }
                    },
                    {
                        "description": "Course ID",
                        "name": "course_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            }
        },
        "/announcements/{id}": {
            "get": {
                "description": "Get an announcement by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "announcements"
                ],
                "summary": "Get an announcement by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Announcement ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an announcement",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "announcements"
                ],
                "summary": "Update an announcement",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Announcement ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Announcement",
                        "name": "announcement",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Announcement"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an announcement",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "announcements"
                ],
                "summary": "Delete an announcement",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Announcement ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Announcement": {
            "type": "object",
            "required": [
                "announcement_description",
                "created_by"
            ],
            "properties": {
                "announcement_description": {
                    "type": "string"
                },
                "announcement_id": {
                    "type": "integer"
                },
                "created_by": {
                    "type": "integer"
                }
            }
        },
        "utils.ApiResponse": {
            "type": "object",
            "properties": {
                "body": {},
                "message": {
                    "type": "string"
                },
                "meta_data": {},
                "status": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8006",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Fiber API",
	Description:      "This is a sample server for a Fiber API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
