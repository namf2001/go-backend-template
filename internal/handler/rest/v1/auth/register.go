package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/namf2001/go-backend-template/internal/controller/auth"
	"github.com/namf2001/go-backend-template/internal/pkg/response"
	"github.com/namf2001/go-backend-template/internal/pkg/validator"
)

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

// Register handles manual registration
// @Summary      Register user
// @Description  Register a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body auth.RegisterRequest true "Registration info"
// @Success      200  {object} auth.RegisterResponse
// @Failure      400  {object} response.Response
// @Failure      500  {object} response.Response
// @Router       /auth/register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, err)
		return
	}

	if err := validator.Validate(req); err != nil {
		validationErrors := validator.ValidationErrors(err)
		response.Error(w, fmt.Errorf("validation error: %v", validationErrors))
		return
	}

	input := auth.RegisterInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := h.ctrl.Register(r.Context(), input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, RegisterResponse{Token: token})
}
