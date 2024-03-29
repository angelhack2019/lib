package utility

import (
	"encoding/json"
	"net/http"
)

func ParseRequestJSON(w http.ResponseWriter, r *http.Request, container interface{}) {
	err := json.NewDecoder(r.Body).Decode(container)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
	}
}

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
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if code != http.StatusOK {
		w.WriteHeader(code)
	}

	w.Write(response)
}
