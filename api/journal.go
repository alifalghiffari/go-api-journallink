package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type JournalErrorResponse struct {
	Error string `json:"error"`
	Message string `json:"message"`
}

type Journal struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Isi    string `json:"isi"`
	Status string `json:"status"`
	DateSubmit string `json:"date_submit"`
}

type JournalSuccessResponse struct {
	Journal []Journal `json:"journal"`
	Message string    `json:"message"`
	Data   interface{} `json:"data"`
}

type DashboardErrorResponse struct {
	Error string `json:"error"`
}

func (api *API) JournalList(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)

	response := JournalSuccessResponse{}
	response.Journal = make([]Journal, 0)

	journal, err := api.journalRepo.FetchJournals()
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

	for _, journal := range journal {
		response.Journal = append(response.Journal, Journal{
			UserID: journal.UserID,
			Isi:    journal.Isi,
			Status: journal.Status,
		})
	}

	encoder.Encode(response)
}

func (api *API) JournalDetailByID(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)

	id, _ := strconv.ParseInt(req.URL.Query().Get("id"), 0, 64)
	journal, err := api.journalRepo.GetJournalByID(id)
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

	journalResponse := JournalSuccessResponse{
		Message: "success",
		Data:     journal,
	}
		
	encoder.Encode(journalResponse)
}

func (api *API) JournalCreate(w http.ResponseWriter, req *http.Request) {
    var journal Journal
	encoder := json.NewEncoder(w)

	user_id, _ := strconv.ParseInt(req.URL.Query().Get("user_id"), 0, 64)
    err := json.NewDecoder(req.Body).Decode(&journal)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(DashboardErrorResponse{Error: err.Error()})
        return
    }

    err = api.journalRepo.InsertJournal(user_id, journal.Isi, journal.Status, journal.DateSubmit)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

	journalResponse := JournalSuccessResponse{
		Message: "create journal success",
		Data:     journal,
	}
		
	encoder.Encode(journalResponse)
}

func (api *API) JournalUpdate(w http.ResponseWriter, req *http.Request) {
	var journal Journal
	encoder := json.NewEncoder(w)

	id, _ := strconv.ParseInt(req.URL.Query().Get("id"), 0, 64)
	err := json.NewDecoder(req.Body).Decode(&journal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(DashboardErrorResponse{Error: err.Error()})
		return
	}

	err = api.journalRepo.UpdateJournal(journal.Isi, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	journalResponse := JournalSuccessResponse{
		Message: "update journal success",
		Data:     journal,
	}
		
	encoder.Encode(journalResponse)
}

// func (api *API) JournalUpdateStatus(w http.ResponseWriter, req *http.Request) {
// 	var journal JournalList
// 	err := json.NewDecoder(req.Body).Decode(&journal)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	err = api.journalRepo.UpdateStatus(journal.Status)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// }