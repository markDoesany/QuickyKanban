package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ProjectID   uint      `gorm:"constraint:OnDelete:CASCADE;" json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"content"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	MaxWorkers  uint      `json:"max_workers"`
	Duedate     time.Time `json:"due_date"`
	Users       []User    `gorm:"many2many:task_users" json:"users"`
	Comments    []Comment `json:"comments"`
}
