package dto

import (
	"fmt"
	"math"
	"nashrul-be/crm/utils/translate"
	"net/http"
)

type BaseResponse struct {
	Code    int               `json:"code,omitempty"`
	Message string            `json:"message,omitempty"`
	Data    any               `json:"data,omitempty"`
	Error   map[string]string `json:"error,omitempty"`
}

func ErrorNotFound(entity string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("%s not found", entity),
	}
}

func ErrorBadRequest(msgErr string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusBadRequest,
		Message: msgErr,
	}
}

func ErrorValidation(err error) BaseResponse {
	return BaseResponse{
		Code:    http.StatusBadRequest,
		Message: "Invalid request parameter",
		Error:   translate.Translate(err),
	}
}

func ErrorInternalServerError() BaseResponse {
	return BaseResponse{
		Code:    http.StatusInternalServerError,
		Message: "Oops, something wrong!",
	}
}

func ErrorUnauthorizedDefault() BaseResponse {
	return ErrorUnauthorized("Unauthorized")
}

func ErrorUnauthorized(msg string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}

func ErrorForbidden() BaseResponse {
	return BaseResponse{
		Code:    http.StatusForbidden,
		Message: "Forbidden",
	}
}

func Authenticated(username, role string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusOK,
		Message: "Authenticated",
		Data: map[string]any{
			"username": username,
			"role":     role,
		},
	}
}

func Success(msg string, data any) BaseResponse {
	return BaseResponse{
		Code:    http.StatusOK,
		Message: msg,
		Data:    data,
	}
}

func SuccessPagination(msg string, currentPage, perpage, total int, data any) BaseResponse {
	return BaseResponse{
		Code:    http.StatusOK,
		Message: msg,
		Data: map[string]any{
			"total":      total,
			"page":       currentPage,
			"perpage":    perpage,
			"total_page": int(math.Ceil(float64(total) / float64(perpage))),
			"data":       data,
		},
	}
}

func Created(msg string, data any) BaseResponse {
	return BaseResponse{
		Code:    http.StatusCreated,
		Message: msg,
		Data:    data,
	}
}
