package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type NotificationErrorResponse struct {
	Error string `json:"error"`
	Message string `json:"message"`
}

type Notification struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Message string `json:"message"`
	Status string `json:"status"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type NotificationSuccessResponse struct {
	Notification []Notification `json:"notification"`
}

func (api *API) GetNotification(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)

	response := NotificationSuccessResponse{}
	response.Notification = make([]Notification, 0)

	user_id, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(NotificationErrorResponse{Error: err.Error()})
		return
	}

	notification, err := api.notificationRepo.GetNotificationByUserID(user_id)
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(NotificationErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	response.Notification = append(response.Notification, Notification{
		ID:     notification.ID,
		UserID: notification.UserID,
		Message: notification.Message,
	})

	encoder.Encode(response)
}

func (api *API) SendNotification(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)

	notification := Notification{}
	err := json.NewDecoder(req.Body).Decode(&notification)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(NotificationErrorResponse{Error: err.Error()})
		return
	}

	err = api.notificationRepo.SendNotification(notification.ID)
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(NotificationErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	response := NotificationSuccessResponse{}
	response.Notification = make([]Notification, 0)
	
	response.Notification = append(response.Notification, Notification{
		Message : "Journal dengan status" + notification.Status + "pada tanggal" + notification.Created_at.Format("2006-01-02 15:04:05") + "telah dikirimkan",
	})

	encoder.Encode(response)
}