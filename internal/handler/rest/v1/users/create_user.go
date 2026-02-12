package users

import (
	"net/http"

	ctrlUsers "github.com/namf2001/go-backend-template/internal/controller/users"
	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/namf2001/go-backend-template/internal/pkg/httpserv"
	"github.com/namf2001/go-backend-template/internal/pkg/validator"
)

// CreateUserRequest represents the request for creating a user
type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required,min=2,max=100"`
}

// CreateUserResponse represents the response for creating a user
type CreateUserResponse struct {
	User model.User `json:"user"`
}

// CreateUser handles the creation of a new user
// @Summary      Create user
// @Description  Create a new user account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        input body users.CreateUserRequest true "User info"
// @Success      201  {object} users.CreateUserResponse
// @Failure      400  {object} httpserv.Error
// @Failure      409  {object} httpserv.Error
// @Failure      500  {object} httpserv.Error
// @Security     BearerAuth
// @Router       /users [post]
func (h Handler) CreateUser() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		var req CreateUserRequest
		if err := httpserv.ParseJSON(r.Body, &req); err != nil {
			return err
		}

		if err := validator.Validate(req); err != nil {
			return webErrValidationFailed
		}

		input := ctrlUsers.CreateUserInput{
			Email: req.Email,
			Name:  req.Name,
		}

		user, err := h.userCtrl.CreateUser(r.Context(), input)
		if err != nil {
			return convertError(err)
		}

		w.WriteHeader(http.StatusCreated)
		httpserv.RespondJSON(r.Context(), w, CreateUserResponse{User: user})
		return nil
	})
}
