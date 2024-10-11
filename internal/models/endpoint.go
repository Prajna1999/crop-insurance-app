package models

import "gorm.io/gorm"

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
	DELETE Method = "DELETE"
	HEAD   Method = "HEAD"
	TRACE  Method = "TRACE"
)

type EndPoint struct {
	gorm.Model
	ApiID       uint
	Path        string `json:"path"`
	Method      Method `json:"method"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Parameters  []Parameter
}
