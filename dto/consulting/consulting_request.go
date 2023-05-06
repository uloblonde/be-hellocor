package consultingdto

import "time"

type CreateConsultingRequest struct {
	Id             int       `json:"id"`
	UserId         int       `json:"userId" validate:"required"`
	FullName       string    `json:"fullName" gorm:"type:varchar(255)"`
	Phone          string    `json:"phone" gorm:"type:varchar(255)"`
	BornDate       time.Time `json:"bornDate" `
	Age            int       `json:"age" `
	Height         int       `json:"height" `
	Weight         int       `json:"weight" `
	Subject        string    `json:"subject" `
	LiveConsulting time.Time `json:"liveConsul"`
	Description    string    `json:"description" `
}
