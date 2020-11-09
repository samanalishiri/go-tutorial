package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

type jsonResponse map[string]interface{}

func postError(response http.ResponseWriter, code int) {
	http.Error(response, http.StatusText(code), code)
}

func postBodyResponse(response http.ResponseWriter, code int, content jsonResponse) {
	if content != nil {
		body, err := json.Marshal(content)
		if err != nil {
			postError(response, http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(code)
		response.Write(body)
		return
	}
	response.WriteHeader(code)
	response.Write([]byte(http.StatusText(code)))
}

func postOptionsResponse(w http.ResponseWriter, methods []string, content jsonResponse) {
	w.Header().Set("Allow", strings.Join(methods, ","))
	postBodyResponse(w, http.StatusOK, content)
}