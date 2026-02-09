package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/namf2001/go-backend-template/internal/controller/auth"
	"github.com/namf2001/go-backend-template/internal/pkg/response"
	"github.com/namf2001/go-backend-template/internal/pkg/validator"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Login handles manual login
// @Summary      User login
// @Description  Authenticate user and return token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body auth.LoginRequest true "Login credentials"
// @Success      200  {object} auth.LoginResponse
// @Failure      400  {object} response.Response
// @Failure      401  {object} response.Response
// @Failure      500  {object} response.Response
// @Router       /auth/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, err)
		return
	}

	if err := validator.Validate(req); err != nil {
		validationErrors := validator.ValidationErrors(err)
		response.Error(w, fmt.Errorf("validation error: %v", validationErrors))
		return
	}

	input := auth.ValidationInput{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := h.ctrl.Login(r.Context(), input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, LoginResponse{Token: token})
}
