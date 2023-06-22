package configuration

type SessionDurationRequest struct {
	Duration string `json:"duration" binding:"required,numeric,gt=0"`
}
