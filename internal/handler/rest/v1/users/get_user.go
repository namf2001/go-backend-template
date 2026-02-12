package users

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/namf2001/go-backend-template/internal/pkg/httpserv"
)

// GetUserResponse represents the response for getting a user
type GetUserResponse struct {
	User model.User `json:"user"`
}

// GetUser handles the retrieval of a user by ID
// @Summary      Get user
// @Description  Get user details by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object} users.GetUserResponse
// @Failure      400  {object} httpserv.Error
// @Failure      404  {object} httpserv.Error
// @Failure      500  {object} httpserv.Error
// @Security     BearerAuth
// @Router       /users/{id} [get]
func (h Handler) GetUser() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return webErrInvalidID
		}

		user, err := h.userCtrl.GetUser(r.Context(), id)
		if err != nil {
			return convertError(err)
		}

		httpserv.RespondJSON(r.Context(), w, GetUserResponse{User: user})
		return nil
	})
}
