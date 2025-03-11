package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `json:"text"`
	UserID  uint   `json:"user_id"`
	TaskID  *uint  `json:"task_id"`
}
