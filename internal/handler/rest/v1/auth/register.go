package auth

import (
	"net/http"

	ctrlAuth "github.com/namf2001/go-backend-template/internal/controller/auth"
	"github.com/namf2001/go-backend-template/internal/pkg/httpserv"
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
// @Failure      400  {object} httpserv.Error
// @Failure      500  {object} httpserv.Error
// @Router       /auth/register [post]
func (h *Handler) Register() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		var req RegisterRequest
		if err := httpserv.ParseJSON(r.Body, &req); err != nil {
			return err
		}

		if err := validator.Validate(req); err != nil {
			return webErrValidationFailed
		}

		input := ctrlAuth.RegisterInput{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		}

		token, err := h.ctrl.Register(r.Context(), input)
		if err != nil {
			return err
		}

		httpserv.RespondJSON(r.Context(), w, RegisterResponse{Token: token})
		return nil
	})
}
