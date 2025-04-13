package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Task struct {
	ID          uint       `json:"id" gorm:"primary_key;auto_increment"`
	Title       string     `json:"title" gorm:"type:varchar(255);not null"`
	Description string     `json:"description" gorm:"type:text;not null"`
	Status      string     `json:"status" binding:"required,oneof=Pending In-Progress Completed"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

func (t *Task) BeforeCreate(scope *gorm.Scope) error {
	log.Printf(" Creating task: %+v", t)
	return nil
}
