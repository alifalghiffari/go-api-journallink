package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Account struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type AccountErrorResponse struct {
	Error string `json:"error"`
}

type AccountSuccessResponse struct {
	User    []User      `json:"user"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (api *API) UserDetailbyID(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)

	id, _ := strconv.ParseInt(req.URL.Query().Get("id"), 0, 64)
	user, err := api.usersRepo.GetUserByID(id)
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(DashboardErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	response := AccountSuccessResponse{}
	response.User = make([]User, 0)
	
	response.User = append(response.User, User{
		Username: user.Username,
		Email:    user.Email,
	})

	encoder.Encode(response)
}