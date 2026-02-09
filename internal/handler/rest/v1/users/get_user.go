package users

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/namf2001/go-backend-template/internal/pkg/response"
)

// GetUser handles the retrieval of a user by ID
// @Summary      Get user
// @Description  Get user details by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object} model.User
// @Failure      400  {object} response.Response
// @Failure      404  {object} response.Response
// @Failure      500  {object} response.Response
// @Router       /users/{id} [get]
func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(w, err)
		return
	}

	user, err := h.userCtrl.GetUser(r.Context(), id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, user)
}
