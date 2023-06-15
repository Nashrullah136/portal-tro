package actor

import (
	"nashrul-be/crm/dto"
	"time"
)

//TODO: create custom binding rule for username and password

func actorNotFound() dto.BaseResponse {
	return dto.ErrorNotFound("Actor")
}

type CreateRequest struct {
	Name     string `json:"name" binding:"omitempty,printascii"`
	Username string `json:"username" binding:"required,printascii"`
	Password string `json:"password" binding:"required,printascii"`
}

type UpdateRequest struct {
	Username string `uri:"username" binding:"required,printascii"`
	Name     string `json:"name" binding:"omitempty,printascii"`
	Password string `json:"password" binding:"omitempty,printascii"`
}

type ChangeActiveRequest struct {
	Activate   []string `json:"activate"`
	Deactivate []string `json:"deactivate"`
}

type ChangePasswordRequest struct {
	Username    string `json:"-"`
	OldPassword string `json:"old_password" binding:"required,printascii"`
	Password    string `json:"password" binding:"required,printascii"`
}

type Representation struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

type PaginationRequest struct {
	PerPage  uint   `form:"perpage" binding:"numeric,gt=0"`
	Page     uint   `form:"page" binding:"numeric,gt=0"`
	Username string `form:"username" binding:"omitempty,printascii"`
	Role     string `form:"role" binding:"omitempty,printascii"`
}
