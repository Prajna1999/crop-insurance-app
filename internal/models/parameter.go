package models

import (
	"gorm.io/gorm"
)

type Inlocation string

type DataType string

const (
	QUERY  Inlocation = "QUERY"
	PATH   Inlocation = "PATH"
	HEADER Inlocation = "HEADER"
	BODY   Inlocation = "BODY"
)

const (
	STRING   DataType = "STRING"
	INTEGER  DataType = "INTEGER"
	NUMBER   DataType = "NUMBER"
	BOOLEAN  DataType = "BOOLEAN"
	ARRAY    DataType = "ARRAY"
	OBJECT   DataType = "OBJECT" //Example: {"name": "John", "age": 30}
	FILE     DataType = "FILE"
	DATE     DataType = "DATE"     //Example: "2023-10-11"
	DATETIME DataType = "DATETIME" //Example: "2023-10-11T14:30:00Z"
	UUID     DataType = "UUID"
	EMAIL    DataType = "EMAIL"
	URI      DataType = "URI" //Example: "https://www.example.com"
	HOSTNAME DataType = "HOSTNAME"
	IPV4     DataType = "IPV4"
	IPV6     DataType = "IPV6"
)

type Parameter struct {
	gorm.Model
	EndPointID  uint       `json:"endpoint_id"`
	Name        string     `json:"name"`
	Required    bool       `json:"required" default:"false"`
	Description string     `json:"description"`
	Inlocation  Inlocation `json:"in_location"`
	DataType    DataType   `json:"data_type" default:"STRING"`
}
