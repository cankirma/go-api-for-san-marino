// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "your@mail.com"
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
        "/v1/category": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update Category.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "update Category",
                "parameters": [
                    {
                        "description": "CategoryId",
                        "name": "CategoryId",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    },
                    {
                        "description": "CategoryName",
                        "name": "CategoryName",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    },
                    {
                        "description": "Description",
                        "name": "Description",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    },
                    {
                        "description": "Picture",
                        "name": "Picture",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new category.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "creates a new category",
                "parameters": [
                    {
                        "description": "CategoryId",
                        "name": "CategoryId",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    },
                    {
                        "description": "CategoryName",
                        "name": "CategoryName",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    },
                    {
                        "description": "Description",
                        "name": "Description",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    },
                    {
                        "description": "Picture",
                        "name": "Picture",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Category"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Category  by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Category  by given ID",
                "parameters": [
                    {
                        "description": "category ID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/category/getall": {
            "get": {
                "description": "Get all exists category.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "gets all exists category",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Category"
                            }
                        }
                    }
                }
            }
        },
        "/v1/category/{id}": {
            "get": {
                "description": "Get Category by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "get Category by given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "CategoryID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Category"
                        }
                    }
                }
            }
        },
        "/v1/token/renew": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Renew access and refresh tokens.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "renew access and refresh tokens",
                "parameters": [
                    {
                        "description": "Refresh token",
                        "name": "refresh_token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user/sign/in": {
            "post": {
                "description": "Auth user and return access and refresh token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "auth user and return access and refresh token",
                "parameters": [
                    {
                        "description": "User Email",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    },
                    {
                        "description": "User Password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user/sign/out": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "De-authorize user and delete refresh token from Redis.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "de-authorize user and delete refresh token from Redis",
                "responses": {
                    "204": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user/sign/up": {
            "post": {
                "description": "Create a new user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "creates a new user",
                "parameters": [
                    {
                        "description": "Email",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    },
                    {
                        "description": "Password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    },
                    {
                        "description": "User role",
                        "name": "user_role",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "primitive"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.Category": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "category_name": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                }
            }
        },
        "entities.User": {
            "type": "object",
            "required": [
                "email",
                "id",
                "password_hash",
                "user_role",
                "user_status"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "password_hash": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_role": {
                    "type": "string"
                },
                "user_status": {
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

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "API",
	Description: "This is an auto-generated API Docs.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
