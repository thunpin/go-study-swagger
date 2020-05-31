// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "HTTP server in Go with Swagger endpoints definition",
    "title": "go-swagger",
    "version": "0.1.0"
  },
  "paths": {
    "/healthz": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "checkHealth",
        "responses": {
          "200": {
            "description": "OK message",
            "schema": {
              "type": "string",
              "enum": [
                "OK"
              ]
            }
          }
        }
      }
    },
    "/hello/{user}": {
      "get": {
        "description": "Returns a greeting to the user!",
        "parameters": [
          {
            "type": "string",
            "description": "The name of the user to greet.",
            "name": "user",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the greeting.",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Invalid characters in \"user\" were provided."
          }
        }
      }
    },
    "/oops/{user}": {
      "get": {
        "description": "Returns a greeting to the user!",
        "parameters": [
          {
            "type": "string",
            "description": "The name of the user to greet.",
            "name": "user",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the greeting.",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Invalid characters in \"user\" were provided."
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "HTTP server in Go with Swagger endpoints definition",
    "title": "go-swagger",
    "version": "0.1.0"
  },
  "paths": {
    "/healthz": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "checkHealth",
        "responses": {
          "200": {
            "description": "OK message",
            "schema": {
              "type": "string",
              "enum": [
                "OK"
              ]
            }
          }
        }
      }
    },
    "/hello/{user}": {
      "get": {
        "description": "Returns a greeting to the user!",
        "parameters": [
          {
            "type": "string",
            "description": "The name of the user to greet.",
            "name": "user",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the greeting.",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Invalid characters in \"user\" were provided."
          }
        }
      }
    },
    "/oops/{user}": {
      "get": {
        "description": "Returns a greeting to the user!",
        "parameters": [
          {
            "type": "string",
            "description": "The name of the user to greet.",
            "name": "user",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the greeting.",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Invalid characters in \"user\" were provided."
          }
        }
      }
    }
  }
}`))
}
