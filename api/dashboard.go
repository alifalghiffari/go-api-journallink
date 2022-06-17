package api

type DashboardErrorResponse struct {
	Error string `json:"error"`
}

type journal struct {
	ID 	   int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Isi    string `json:"isi"`
	Status string `json:"status"`
	DateSubmit string `json:"date_submit"`
}

type DashboardSuccessResponse struct {
	Journal []JournalList `json:"journal"`
}
