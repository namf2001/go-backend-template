package users

import (
	"net/http"
	"strconv"

	ctrlUsers "github.com/namf2001/go-backend-template/internal/controller/users"
	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/namf2001/go-backend-template/internal/pkg/httpserv"
)

// ListUsersRequest represents the request for listing users
type ListUsersRequest struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Email  string `json:"email"`
}

// ListUsersResponse represents the response for listing users
type ListUsersResponse struct {
	Users  []model.User `json:"users"`
	Total  int64        `json:"total"`
	Limit  int          `json:"limit"`
	Offset int          `json:"offset"`
}

// ListUsers handles the listing of users with optional filters
// @Summary      List users
// @Description  Get a list of users
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        limit  query     int     false  "Limit"
// @Param        offset query     int     false  "Offset"
// @Param        email  query     string  false  "Email filter"
// @Success      200  {object} users.ListUsersResponse
// @Failure      500  {object} httpserv.Error
// @Security     BearerAuth
// @Router       /users [get]
func (h Handler) ListUsers() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		// Parse query parameters
		limitStr := r.URL.Query().Get("limit")
		offsetStr := r.URL.Query().Get("offset")
		email := r.URL.Query().Get("email")

		limit := 10 // default
		if limitStr != "" {
			if l, err := strconv.Atoi(limitStr); err == nil {
				limit = l
			}
		}

		offset := 0
		if offsetStr != "" {
			if o, err := strconv.Atoi(offsetStr); err == nil {
				offset = o
			}
		}

		filters := ctrlUsers.ListFilters{
			Limit:  limit,
			Offset: offset,
			Email:  email,
		}

		result, totalUser, err := h.userCtrl.ListUsers(r.Context(), filters)
		if err != nil {
			return convertError(err)
		}

		httpserv.RespondJSON(r.Context(), w, ListUsersResponse{
			Users:  result,
			Total:  totalUser,
			Limit:  limit,
			Offset: offset,
		})
		return nil
	})
}
