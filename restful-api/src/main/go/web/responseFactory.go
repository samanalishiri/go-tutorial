package web

import (
	"encoding/json"
	"net/http"
	"strings"
)

type JsonResponse map[string]interface{}

func ThrowError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func CreatHttpResponse(w http.ResponseWriter, code int, content JsonResponse) {
	if content == nil {
		w.WriteHeader(code)
		w.Write([]byte(http.StatusText(code)))
		return
	}

	body, err := json.Marshal(content)
	if err != nil {
		ThrowError(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(body)
}

func CreateOptionsResponse(w http.ResponseWriter, methods []string, content JsonResponse) {
	w.Header().Set("Allow", strings.Join(methods, ","))
	CreatHttpResponse(w, http.StatusOK, content)
}
