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
    }
  },
  "components": {
    "schemas": {
      "hospital": {
        "type": "object",
        "properties": {
          "HID": {
            "type": "string",
            "example": "0001"
          },
          "Title": {
            "type": "string",
            "example": "name hospital"
          }
        },
        "xml": {
          "name": "hospital"
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
	Title:            "Party's Applicant",
	Description:      "MVC Example 01",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
