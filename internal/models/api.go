package models

import "gorm.io/gorm"

type Api struct {
	gorm.Model
	Name           string `json:"name"`
	Description    string `json:"description"`
	BaseUrl        string `json:"base_url"`
	Version        string `json:"version" default:"1"`
	CategoryID     uint   `json:"category_id"`
	OrganizationID uint   `json:"organization_id"`
	IsPublic       bool   `json:"is_public" default:"false"`
	EndPoints      []EndPoint
}
