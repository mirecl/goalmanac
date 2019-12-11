package http

import (
	"context"
	"encoding/json"
	"net/http"
)

type cntHTTPEvent struct {
	Result int `json:"result"`
}

// cntHandler - handler для пути /api/count_event
func (api *APIServerHTTP) cntHandler(w http.ResponseWriter, r *http.Request) {
	cnt, err := api.Event.GetCount(context.Background())
	if err != nil {
		api.Logger.Errorf("%s", err.Error())
		api.Helper.Error(w, err, http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cntHTTPEvent{Result: *cnt})
}
