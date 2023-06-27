package span

type UpdateRequest struct {
	DocumentNumber string `json:"documentNumber" binding:"required"`
}
