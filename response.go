package main

import (
	"encoding/json"
	"net/http"
)

// function untuk membaca json body ke struct tujuan
func decodeJSON(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(dst)
}

// function untuk mengirim response json
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

// function untuk mengirim response error
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorResponse{
		Error: message,
	})
}
