package users

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/namf2001/go-backend-template/internal/pkg/response"
)

// DeleteUser handles the deletion of a user by ID
// @Summary      Delete user
// @Description  Delete a user account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      204  {object} nil
// @Failure      400  {object} response.Response
// @Failure      500  {object} response.Response
// @Router       /users/{id} [delete]
func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(w, err)
		return
	}

	if err := h.userCtrl.DeleteUser(r.Context(), id); err != nil {
		response.Error(w, err)
		return
	}

	response.NoContent(w)
}
