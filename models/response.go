package models

import "time"

type Response struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	UserId       int
	User         User
	ConsulId     int
	ResponseText string `json:"responseText" gorm:"type: varchar(255)" `
	ConsulLink   string `json:"consulLink" gorm:"type: varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
