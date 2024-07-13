package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ValuesRequest struct {
	Values []int `json:"values"`
}

func CalculateHighestHandler(w http.ResponseWriter, r *http.Request) {

	var req ValuesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var highest int
	for _, v := range req.Values {
		if v > highest {
			highest = v
		}
	}

	if highest == 50 {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "highest number is too big"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "%d", highest)
}
