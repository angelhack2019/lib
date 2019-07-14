package utility

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func Respond(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"response": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if code != http.StatusOK {
		w.WriteHeader(code)
	}

	w.Write(response)
}
