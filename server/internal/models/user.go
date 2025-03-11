package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `gorm:"unique; not null" json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	Project  []Project `gorm:"many2many:project_users" json:"projects"`
	Tasks    []Task    `gorm:"many2many:task_users;" json:"tasks"`
	Comments []Comment `json:"comments"`
	ImageURL string    `json:"image_url"`
}
