package span

type UpdateRequest struct {
	DocumentNumber string `json:"documentNumber" binding:"required"`
}

type Presentation struct {
	DocumentNumber      string
	DocumentDate        string
	BeneficiaryBankCode string
	StatusCode          string
	EmailAddress        string
	BeneficiaryAccount  string
	Amount              string
	BeneficiaryBank     string
	IsPatched           bool `json:"is_patched"`
}
