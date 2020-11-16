package user

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

type JsonResponse map[string]interface{}

func ThrowError(c echo.Context, code int, meg string) {
	c.String(code, meg)
}

func CreatHttpResponse(c echo.Context, code int, content interface{}) {
	if content == nil {
		c.String(code, "")
		return
	}

	body, err := json.Marshal(content)
	if err != nil {
		ThrowError(c, http.StatusInternalServerError, "")
		return
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(code)
	c.Response().Write(body)
}

func CreateOptionsResponse(c echo.Context, methods []string, content JsonResponse) {
	c.Response().Header().Set("Allow", strings.Join(methods, ","))
	CreatHttpResponse(c, http.StatusOK, content)
}
