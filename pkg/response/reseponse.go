package response

import (
	"github.com/linskybing/platform-go/internal/domain/group"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type TokenResponse struct {
	Token    string `json:"token"`
	UID      uint   `json:"user_id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_super_admin"`
}

type GroupResponse struct {
	Message string      `json:"message"`
	Group   group.Group `json:"group"`
}

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
