package entities

import (
	"time"
)

type Task struct {
	ID              uint      `gorm:"primaryKey; autoIncrement" json:"id"`
	TaskName        string    `json:"task_name"`
	TaskDescription string    `json:"task_description"`
	CreatedAt       time.Time `gorm:"autoCreatedTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdatedTime" json:"updated_at"`
}
