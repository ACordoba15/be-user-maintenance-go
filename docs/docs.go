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
        "/record": {
            "post": {
                "description": "Agrega un nuevo registro a la base de datos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "record"
                ],
                "summary": "Crea un nuevo registro",
                "parameters": [
                    {
                        "description": "Información del nuevo registro",
                        "name": "record",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Record"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Record"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/record/all": {
            "get": {
                "description": "Retorna una lista de todos los registros almacenados en la base de datos.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "record"
                ],
                "summary": "Obtiene todos los registros",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Record"
                            }
                        }
                    }
                }
            }
        },
        "/record/{id}": {
            "get": {
                "description": "Retorna un registro específico basado en el ID proporcionado.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "record"
                ],
                "summary": "Obtiene un registro por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del registro",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Record"
                        }
                    },
                    "404": {
                        "description": "Record Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Actualiza la información de un registro.",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "record"
                ],
                "summary": "Actualiza un registro",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Record"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Realiza el borrado lógico de un registro específico.",
                "tags": [
                    "record"
                ],
                "summary": "Elimina un registro por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del registro",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "No"
                        }
                    },
                    "404": {
                        "description": "Record Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "put": {
                "description": "Actualiza la información de un registro.",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Actualiza un registro",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Agrega un nuevo registro a la base de datos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Crea un nuevo registro",
                "parameters": [
                    {
                        "description": "Información del nuevo registro",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Record"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/all": {
            "get": {
                "description": "Retorna una lista de todos los registros almacenados en la base de datos.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Obtiene todos los registros",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Valida un usuari registrado.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login de usuario",
                "parameters": [
                    {
                        "description": "Información del nuevo registro",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Record"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Retorna un registro específico basado en el ID proporcionado.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Obtiene un registro por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del registro",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "404": {
                        "description": "Record Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Realiza el borrado lógico de un registro específico.",
                "tags": [
                    "user"
                ],
                "summary": "Elimina un registro por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del registro",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "No"
                        }
                    },
                    "404": {
                        "description": "User Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Record": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api/",
	Schemes:          []string{},
	Title:            "API de ejemplo con Swagger y Gorilla Mux",
	Description:      "Esta es una API en Go documentada con Swagger y usando Gorilla Mux.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
