package users

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	ctrlUsers "github.com/namf2001/go-backend-template/internal/controller/users"
	"github.com/namf2001/go-backend-template/internal/pkg/httpserv"
	"github.com/namf2001/go-backend-template/internal/pkg/validator"
)

// UpdateUserRequest represents the request for updating a user
type UpdateUserRequest struct {
	Email string `json:"email" validate:"omitempty,email"`
	Name  string `json:"name" validate:"omitempty,min=2,max=100"`
}

// UpdateUser handles the updating of a user by ID
// @Summary      Update user
// @Description  Update user details
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id     path      int                    true  "User ID"
// @Param        input  body      users.UpdateUserRequest  true  "Update info"
// @Success      204  {object} nil
// @Failure      400  {object} httpserv.Error
// @Failure      500  {object} httpserv.Error
// @Security     BearerAuth
// @Router       /users/{id} [put]
func (h Handler) UpdateUser() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return webErrInvalidID
		}

		var req UpdateUserRequest
		if err := httpserv.ParseJSON(r.Body, &req); err != nil {
			return err
		}

		if err := validator.Validate(req); err != nil {
			return webErrValidationFailed
		}

		input := ctrlUsers.UpdateUserInput{
			Email: req.Email,
			Name:  req.Name,
		}

		if err := h.userCtrl.UpdateUser(r.Context(), id, input); err != nil {
			return convertError(err)
		}

		w.WriteHeader(http.StatusNoContent)
		return nil
	})
}
