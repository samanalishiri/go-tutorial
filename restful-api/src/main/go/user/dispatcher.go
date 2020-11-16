package user

import (
	"net/http"
)

var (
	handler = NewHandler()
)

type Dispatcher struct {
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

func (d *Dispatcher) Init() FrontController {
	controller := NewFrontController()

	controller.Get("/users", handler.GetAll)
	controller.Post("/users", handler.SaveOne)
	controller.Get("/users", handler.GetAll)
	controller.Options("/users", func(c Context) {
		CreateOptionsResponse(c,
			[]string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions},
			nil)
	})
	controller.Get("/users/:id", handler.GetOne)
	controller.Put("/users/:id", handler.UpdateOne)
	controller.Patch("/users/:id", handler.PatchOne)
	controller.Delete("/users/:id", handler.DeleteOne)
	controller.Head("/users/:id", handler.GetOne)
	controller.Options("/users/:id", func(c Context) {
		CreateOptionsResponse(c,
			[]string{
				http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions},
			nil)
	})
	return controller
}
