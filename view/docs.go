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
    },
    {
      "name": "count",
      "description": "Count Patient"
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
    },
    "/listPatientStatus": {
      "get": {
        "tags": [
          "patient"
        ],
        "summary": "Get Patient Positive",
        "operationId": "getPatientStatus",
        "responses": {
          "200": {
            "description": "successful operation",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/patientStatus"
                }
              },
              "application/xml": {
                "schema": {
                  "$ref": "#/components/schemas/patientStatus"
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
    "/countPatient": {
      "get": {
        "tags": [
          "count"
        ],
        "summary": "Count Patient Positive",
        "operationId": "countPatient",
        "responses": {
          "200": {
            "description": "successful operation",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/count"
                }
              },
              "application/xml": {
                "schema": {
                  "$ref": "#/components/schemas/count"
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
    "/countPerHospital": {
      "get": {
        "tags": [
          "hospital"
        ],
        "summary": "Count Patient Positive Per Hospital",
        "operationId": "countHospital",
        "responses": {
          "200": {
            "description": "successful operation",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/countHospital"
                }
              },
              "application/xml": {
                "schema": {
                  "$ref": "#/components/schemas/countHospital"
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
          "HNID": {
            "type": "string",
            "example": "0001"
          },
          "Firstname": {
            "type": "string",
            "example": "Firstname"
          },
          "Lastname": {
            "type": "string",
            "example": "Lastname"
          }
        },
        "xml": {
          "name": "patient"
        }
      },
      "count": {
        "type": "object",
        "properties": {
          "CountPatient": {
            "type": "integer",
            "example": 1
          }
        },
        "xml": {
          "name": "count"
        }
      },
      "patientStatus": {
        "type": "object",
        "properties": {
          "HNID": {
            "type": "string",
            "example": "0001"
          },
          "Firstname": {
            "type": "string",
            "example": "Firstname"
          },
          "Lastname": {
            "type": "string",
            "example": "Lastname"
          },
          "CountStatus": {
            "type": "string",
            "example": "Positive/Negative"
          }
        },
        "xml": {
          "name": "patientStatus"
        }
      },
      "countHospital": {
        "type": "object",
        "properties": {
          "HID": {
            "type": "string",
            "example": "0001"
          },
          "Title": {
            "type": "string",
            "example": "Hospital name"
          },
          "CountPatient": {
            "type": "integer",
            "example": 1
          }
        },
        "xml": {
          "name": "countHospital"
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
