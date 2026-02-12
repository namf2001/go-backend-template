package users

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/namf2001/go-backend-template/internal/pkg/httpserv"
)

// DeleteUser handles the deletion of a user by ID
// @Summary      Delete user
// @Description  Delete a user account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      204  {object} nil
// @Failure      400  {object} httpserv.Error
// @Failure      500  {object} httpserv.Error
// @Security     BearerAuth
// @Router       /users/{id} [delete]
func (h Handler) DeleteUser() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return webErrInvalidID
		}

		if err := h.userCtrl.DeleteUser(r.Context(), id); err != nil {
			return convertError(err)
		}

		w.WriteHeader(http.StatusNoContent)
		return nil
	})
}
