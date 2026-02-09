package auth

import (
	"github.com/namf2001/go-backend-template/internal/controller/auth"
)

type Handler struct {
	ctrl auth.Controller
}

func New(ctrl auth.Controller) *Handler {
	return &Handler{
		ctrl: ctrl,
	}
}
