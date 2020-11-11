package root

import (
	"net/http"
)

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("The URL is invalid"))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("GO Tutorial"))
}
