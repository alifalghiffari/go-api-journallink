package api

import (
	"encoding/json"
	"net/http"
)

type JournalListErrorResponse struct {
	Error string `json:"error"`
}

type JournalList struct {
	ID        int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Isi    string `json:"isi"`
	Status string `json:"status"`
	DateSubmit string `json:"date_submit"`
}

type JournalListSuccessResponse struct {
	Journal []JournalList `json:"journal"`
}

func (api *API) JournalList(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)

	response := JournalListSuccessResponse{}
	response.Journal = make([]JournalList, 0)

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
		response.Journal = append(response.Journal, JournalList{
			UserID: journal.UserID,
			Isi:    journal.Isi,
			Status: journal.Status,
		})
	}

	encoder.Encode(response)
}

func (api *API) JournalCreate(w http.ResponseWriter, req *http.Request) {
    var journal JournalList
    err := json.NewDecoder(req.Body).Decode(&journal)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    err = api.journalRepo.InsertJournal(journal.UserID, journal.Isi, journal.Status, journal.DateSubmit)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
}

func (api *API) JournalUpdate(w http.ResponseWriter, req *http.Request) {
	var journal JournalList
	err := json.NewDecoder(req.Body).Decode(&journal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = api.journalRepo.UpdateJournal(journal.Isi)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}