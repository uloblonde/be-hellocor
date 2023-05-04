package models

import "time"

type Consulting struct {
	Id             uint   `json:"id"`
	BornDate       string `json:"email" gorm:"type:varchar(255)"`
	User           UserConsulResponse
	UserId         int    `json:"userId"`
	Age            int    `json:"password" `
	Height         int    `json:"height" `
	Weight         int    `json:"weight" `
	Gender         string `json:"gender" gorm:"type:varchar(255)"`
	Subject        string `json:"address" gorm:"type:varchar(255)"`
	LiveConsulting string `json:"liveConsul" gorm:"type:varchar(255)"`
	Description    string `json:"description" gorm:"type:varchar(255)"`
	Status         string `gorm:"type: varchar(100)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
