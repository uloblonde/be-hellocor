package models

import "time"

type Consulting struct {
	Id             uint `json:"id"`
	User           UserConsulResponse
	UserId         int    `json:"userId"`
	BornDate       string `json:"bornDate"`
	Age            int    `json:"age" `
	Height         int    `json:"height" `
	Weight         int    `json:"weight" `
	Subject        string `json:"subject" gorm:"type:varchar(255)"`
	LiveConsulting string `json:"liveConsul"`
	Description    string `json:"description" gorm:"type:varchar(255)"`
	Status         string `gorm:"type: varchar(100)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
