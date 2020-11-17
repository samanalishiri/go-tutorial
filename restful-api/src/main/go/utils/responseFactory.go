package utils

import (
	"encoding/json"
	"net/http"
	"restfull-api/src/main/go/contract"
	"strings"
)

type JsonResponse map[string]interface{}

func ThrowError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func CreatHttpResponse(c contract.Context, code int, content interface{}) {
	if content == nil {
		c.Writer.WriteHeader(code)
		c.Writer.Write([]byte(http.StatusText(code)))
		return
	}

	body, err := json.Marshal(content)
	if err != nil {
		ThrowError(c.Writer, http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	c.Writer.Write(body)
}

func CreateOptionsResponse(c contract.Context, methods []string, content JsonResponse) {
	c.Writer.Header().Set("Allow", strings.Join(methods, ","))
	CreatHttpResponse(c, http.StatusOK, content)
}
