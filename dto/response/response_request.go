package responsedto

type ResponseRequest struct {
	ResponseText string `json:"responseText" gorm:"type:varchar(255)"`
	ConsulLink   string `json:"consulLink" gorm:"type:varchar(255)"`
}
