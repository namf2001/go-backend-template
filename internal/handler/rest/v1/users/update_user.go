package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/namf2001/go-backend-template/internal/controller/users"
	"github.com/namf2001/go-backend-template/internal/pkg/response"
)

// UpdateUser handles the updating of a user by ID
// @Summary      Update user
// @Description  Update user details
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id     path      int                    true  "User ID"
// @Param        input  body      users.UpdateUserInput  true  "Update info"
// @Success      204  {object} nil
// @Failure      400  {object} response.Response
// @Failure      500  {object} response.Response
// @Router       /users/{id} [put]
func (h Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(w, err)
		return
	}

	var input users.UpdateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, err)
		return
	}

	if err := h.userCtrl.UpdateUser(r.Context(), id, input); err != nil {
		response.Error(w, err)
		return
	}

	response.NoContent(w)
}
