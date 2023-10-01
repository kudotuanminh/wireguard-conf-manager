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
        "/apply": {
            "post": {
                "description": "Responds with the request status",
                "tags": [
                    "server"
                ],
                "summary": "Apply server config",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/profiles": {
            "get": {
                "description": "Responds with the list of all client profiles as JSON",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profiles"
                ],
                "summary": "List all profiles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/main.Client"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/profiles/new": {
            "post": {
                "description": "Responds with the newly created client profile as JSON",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profiles"
                ],
                "summary": "Create new profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Profile Name",
                        "name": "profileName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Client IP",
                        "name": "clientIP",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Allowed IPs",
                        "name": "allowedIPs",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Client"
                        }
                    }
                }
            }
        },
        "/profiles/{profileName}": {
            "get": {
                "description": "Responds with the client profile as JSON",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profiles"
                ],
                "summary": "List profile by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Profile Name",
                        "name": "profileName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Client"
                        }
                    }
                }
            },
            "delete": {
                "description": "Responds with the deleted client profile as JSON",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profiles"
                ],
                "summary": "Delete profile by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Profile Name",
                        "name": "profileName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Client"
                        }
                    }
                }
            }
        },
        "/profiles/{profileName}/getconf": {
            "get": {
                "description": "Responds with the client profile config as plain text",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "profiles"
                ],
                "summary": "Get profile config",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Profile Name",
                        "name": "profileName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Client": {
            "type": "object",
            "properties": {
                "AllowedIPs": {
                    "type": "string"
                },
                "ClientIP": {
                    "type": "string"
                },
                "PrivateKey": {
                    "type": "string"
                },
                "ProfileID": {
                    "type": "integer"
                },
                "ProfileName": {
                    "type": "string"
                },
                "PublicKey": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "WireGuard Configuration Manager API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
