package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `json:"title"`
	Status      string `json:"status"`
	Description string `json:"description"`
	OwnerID     uint   `json:"owner_id"`
	Users       []User `gorm:"many2many:project_users" json:"users"`
	Tasks       []Task `gorm:"constraint:OnDelete:CASCADE;" json:"tasks"`
}
