package users

import (
	"encoding/json"
	"net/http"

	"github.com/namf2001/go-backend-template/internal/controller/users"
	"github.com/namf2001/go-backend-template/internal/pkg/response"
)

// CreateUserResponse represents the response for creating a user
type CreateUserResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CreateUser handles the creation of a new user
// @Summary      Create user
// @Description  Create a new user account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        input body users.CreateUserInput true "User info"
// @Success      201  {object} users.CreateUserResponse
// @Failure      400  {object} response.Response
// @Failure      500  {object} response.Response
// @Router       /users [post]
func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input users.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, err)
		return
	}

	user, err := h.userCtrl.CreateUser(r.Context(), input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Created(w, user)
}
