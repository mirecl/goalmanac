package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type cntHTTPEvent struct {
	Result int `json:"result"`
}

// cntHandler - handler для пути /api/count_event
func (api *APIServerHTTP) cntHandler(w http.ResponseWriter, r *http.Request) {
	cnt, err := api.Event.GetCount(context.Background())
	if err != nil {
		api.Error(w, fmt.Errorf("Error in %s (%s) %w", GetFunc(), "api.Event.GetCount", err), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cntHTTPEvent{Result: *cnt})
}
