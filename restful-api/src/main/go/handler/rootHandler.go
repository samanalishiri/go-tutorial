package handler

import (
	"net/http"
)

func RootHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("The URL is invalid"))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Hello GO lang\n"))
}
