package models

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	OrgName       string `json:"org_name"`
	AdminEmail    string `json:"admin_email"`
	AdminUserName string `json:"admin_user_name"`
	Apis          []Api
}
