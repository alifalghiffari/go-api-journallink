package api

// import (
// 	"encoding/json"
// 	"go-api-project/repository"
// 	"net/http"
// )

type AdminErrorResponse struct {
	Error string `json:"error"`
}

// type AdminResponse struct {
// 	Journals []repository.Journal `json:"journal"`
// }

// func (api *API) getAdminDashboard(w http.ResponseWriter, req *http.Request) {
// 	api.AllowOrigin(w, req)

// 	encoder := json.NewEncoder(w)
// 	journal, err := api.journalRepo.FetchJournals()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		encoder.Encode(AdminErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	encoder.Encode(journal)

// 	return
// }
