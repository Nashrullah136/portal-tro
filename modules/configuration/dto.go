package configuration

type SessionDurationRequest struct {
	Duration int `json:"duration" binding:"required,gt=60"`
}
