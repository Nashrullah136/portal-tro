package rdn

type UpdateRequest struct {
	Brivano string `uri:"brivano" binding:"required"`
	Active  string `json:"active" binding:"required,len=1"`
}
