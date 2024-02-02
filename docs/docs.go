// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "只因哥",
            "email": "www.758494478@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/hello": {
            "get": {
                "description": "测试一下",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hello world"
                ],
                "summary": "hello world",
                "responses": {
                    "200": {
                        "description": "msg\":\"hello world\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/role/add": {
            "post": {
                "description": "添加角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "创建角色",
                "parameters": [
                    {
                        "description": "Role Information",
                        "name": "roleInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RoleInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RoleInfo"
                        }
                    }
                }
            }
        },
        "/role/delete/{id}": {
            "delete": {
                "description": "删除角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "删除角色",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "删除失败",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/role/detail": {
            "get": {
                "description": "获取角色信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "获取角色信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RoleInfo"
                        }
                    }
                }
            }
        },
        "/role/edit/{id}": {
            "put": {
                "description": "修改角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "修改角色",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Role Information",
                        "name": "roleInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RoleInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RoleInfo"
                        }
                    },
                    "400": {
                        "description": "修改失败",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.RoleInfo": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8093",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "规范地写一个go_demo",
	Description:      "从头开始写一个demo",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
