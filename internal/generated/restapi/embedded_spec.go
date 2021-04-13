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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Простейший API для управляения структурой пользователя User, включает в себя создание/получение/обновление/удаление",
    "title": "Swagger User CRUD",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "fry256@yandex.ru"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "petstore.swagger.io",
  "basePath": "/api/",
  "paths": {
    "/user": {
      "post": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Create user",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        }
      }
    },
    "/user/{uid}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Get user by userID",
        "operationId": "getUserByID",
        "parameters": [
          {
            "type": "string",
            "description": "The userID that needs to be fetched. Use 1 for testing. ",
            "name": "uid",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid userID supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "put": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Updated user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "userID that need to be updated",
            "name": "uid",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid user supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Delete user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "description": "The userID that needs to be deleted",
            "name": "uid",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid userID supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "properties": {
        "birthDate": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "uid": {
          "type": "integer",
          "format": "int64"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Operations about user",
      "name": "user",
      "externalDocs": {
        "description": "Find out more about our store",
        "url": "http://swagger.io"
      }
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Простейший API для управляения структурой пользователя User, включает в себя создание/получение/обновление/удаление",
    "title": "Swagger User CRUD",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "fry256@yandex.ru"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "petstore.swagger.io",
  "basePath": "/api/",
  "paths": {
    "/user": {
      "post": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Create user",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        }
      }
    },
    "/user/{uid}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Get user by userID",
        "operationId": "getUserByID",
        "parameters": [
          {
            "type": "string",
            "description": "The userID that needs to be fetched. Use 1 for testing. ",
            "name": "uid",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid userID supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "put": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Updated user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "userID that need to be updated",
            "name": "uid",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid user supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Delete user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "description": "The userID that needs to be deleted",
            "name": "uid",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid userID supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "properties": {
        "birthDate": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "uid": {
          "type": "integer",
          "format": "int64"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Operations about user",
      "name": "user",
      "externalDocs": {
        "description": "Find out more about our store",
        "url": "http://swagger.io"
      }
    }
  ]
}`))
}
