package consultingdto

type CreateConsultingRequest struct {
	Id             int    `json:"id"`
	UserId         int    `json:"userId" validate:"required"`
	FullName       string `json:"fullName" gorm:"type:varchar(255)"`
	Phone          string `json:"phone" gorm:"type:varchar(255)"`
	BornDate       string `json:"bornDate" gorm:"type:varchar(255)"`
	Age            int    `json:"password" `
	Height         int    `json:"height" `
	Weight         int    `json:"weight" `
	Gender         string `json:"gender" gorm:"type:varchar(255)"`
	Subject        string `json:"subject" gorm:"type:varchar(255)"`
	LiveConsulting string `json:"liveConsul"`
	Description    string `json:"description" gorm:"type:varchar(255)"`
}
