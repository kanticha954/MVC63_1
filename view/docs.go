package view

import "github.com/swaggo/swag"

const docTemplate = `{
  "openapi": "3.0.3",
  "info": {
    "title": "Golang MVC Structure",
    "version": "1.0.11"
  },
    "tags": [
    {
      "name": "hospital",
      "description": "Access to Hospital Information",
      "externalDocs": {
        "description": "Find out more about Hospital",
        "url": "http://swagger.io"
      }
    },
    {
      "name": "patient",
      "description": "About Patient"
    }
  ],
  "paths": {
    "/listHospital": {
      "get": {
        "tags": [
          "hospital"
        ],
        "summary": "Get Hospital Information",
        "operationId": "getHospital",
        "responses": {
          "200": {
            "description": "successful operation",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/hospital"
                }
              },
              "application/xml": {
                "schema": {
                  "$ref": "#/components/schemas/hospital"
                }
              }
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Order not found"
          }
        }
      }
    },
    "/listPatient": {
      "get": {
        "tags": [
          "patient"
        ],
        "summary": "Get Patient Information",
        "operationId": "getPatient",
        "responses": {
          "200": {
            "description": "successful operation",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/patient"
                }
              },
              "application/xml": {
                "schema": {
                  "$ref": "#/components/schemas/patient"
                }
              }
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Order not found"
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "hospital": {
        "type": "object",
        "properties": {
          "hospital_id": {
            "type": "string",
            "example": "0001"
          },
          "title": {
            "type": "string",
            "example": "name hospital"
          }
        },
        "xml": {
          "name": "hospital"
        }
      },
      "patient": {
        "type": "object",
        "properties": {
          "hnid": {
            "type": "string",
            "example": "0001"
          },
          "firstname": {
            "type": "string",
            "example": "Firstname"
          },
          "lastname": {
            "type": "string",
            "example": "Lastname"
          }
        },
        "xml": {
          "name": "patient"
        }
    }
  }
}
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Covid Applicant",
	Description:      "MVC Example 01",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
