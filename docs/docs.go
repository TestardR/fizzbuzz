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
        "contact": {
            "name": "Romain Testard",
            "email": "romain.rtestard@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/fizzbuzz": {
            "get": {
                "description": "Fizzbuzz operation",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fizzbuzz Operation"
                ],
                "summary": "Return fizzbuzz result",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "int1 query parameter",
                        "name": "int1",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "int2 query parameter",
                        "name": "int2",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "limit query parameter",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "str1 query parameter",
                        "name": "str1",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "str2 query parameter",
                        "name": "str2",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/ResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Healthcheck endpoint, to ensure that the service is running.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Health"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    }
                }
            }
        },
        "/stats": {
            "get": {
                "description": "Returns the query that was performed the most and also the number of times",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Statistic"
                ],
                "summary": "Retrieve statistic regarding fizzbuzz request",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ResponseSuccess"
                        }
                    },
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Reset all statistics regarding the queries made in the past",
                "tags": [
                    "Statistic"
                ],
                "summary": "Reset statistics",
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "Health": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "boolean"
                }
            }
        },
        "ResponseError": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Error"
                    }
                }
            }
        },
        "ResponseSuccess": {
            "type": "object",
            "properties": {
                "data": {}
            }
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
	Host:        "localhost:3000",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Fizzbuzz Rest Server",
	Description: "This is a fizzbuz server with a statistic endpoint",
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
	swag.Register("swagger", &s{})
}
