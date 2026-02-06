package users

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/namf2001/go-backend-template/internal/pkg/response"
)

// GetUser handles the retrieval of a user by ID
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
