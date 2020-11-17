package root

import (
	"net/http"
	"restfull-api/src/main/go/contract"
	"restfull-api/src/main/go/utils"
)

type HandlerImpl struct {
}

func NewHandler() contract.Handler {
	return &HandlerImpl{}
}

func notSupport(c contract.Context) {
	c.Writer.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
}

func (h *HandlerImpl) SaveOne(c contract.Context) {
	notSupport(c)
}

func (h *HandlerImpl) GetOne(c contract.Context) {
	notSupport(c)
}

func (h *HandlerImpl) UpdateOne(c contract.Context) {
	notSupport(c)
}

func (h *HandlerImpl) DeleteOne(c contract.Context) {
	notSupport(c)
}

func (h *HandlerImpl) PatchOne(c contract.Context) {
	notSupport(c)
}

func (h *HandlerImpl) GetAll(c contract.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte("GO Tutorial"))
}

func (h *HandlerImpl) GetBasicMethod(c contract.Context) {
	utils.CreateOptionsResponse(c, []string{http.MethodGet, http.MethodOptions}, nil)
}

func (h *HandlerImpl) GetAllMethod(c contract.Context) {
	notSupport(c)
}
