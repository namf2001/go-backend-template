package auth

import (
	"net/http"

	ctrlAuth "github.com/namf2001/go-backend-template/internal/controller/auth"
	"github.com/namf2001/go-backend-template/internal/pkg/httpserv"
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
// @Failure      400  {object} httpserv.Error
// @Failure      401  {object} httpserv.Error
// @Failure      500  {object} httpserv.Error
// @Router       /auth/login [post]
func (h *Handler) Login() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		var req LoginRequest
		if err := httpserv.ParseJSON(r.Body, &req); err != nil {
			return err
		}

		if err := validator.Validate(req); err != nil {
			return webErrValidationFailed
		}

		input := ctrlAuth.ValidationInput{
			Email:    req.Email,
			Password: req.Password,
		}

		token, err := h.ctrl.Login(r.Context(), input)
		if err != nil {
			return webErrInvalidCredentials
		}

		httpserv.RespondJSON(r.Context(), w, LoginResponse{Token: token})
		return nil
	})
}
