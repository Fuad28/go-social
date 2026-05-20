package main

import (
	"net/http"
)

func (s *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	// w.Header().Set("Content-Type", "application/json")

	// err := json.NewEncoder(w).Encode(users)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.WriteHeader(http.StatusOK)
}
