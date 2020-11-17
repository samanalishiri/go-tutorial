package user

import "restfull-api/src/main/go/contract"

type Dispatcher struct {
	handler contract.Handler
}

func NewDispatcher() contract.Dispatcher {
	return &Dispatcher{
		handler: NewHandler(),
	}
}

func (d *Dispatcher) Init() contract.FrontController {
	controller := NewFrontController()

	controller.Get("/users", d.handler.GetAll)
	controller.Post("/users", d.handler.SaveOne)
	controller.Head("/users", d.handler.GetAll)
	controller.Options("/users", d.handler.GetBasicMethod)
	controller.Get("/users/:id", d.handler.GetOne)
	controller.Put("/users/:id", d.handler.UpdateOne)
	controller.Patch("/users/:id", d.handler.PatchOne)
	controller.Delete("/users/:id", d.handler.DeleteOne)
	controller.Head("/users/:id", d.handler.GetOne)
	controller.Options("/users/:id", d.handler.GetAllMethod)

	return controller
}
