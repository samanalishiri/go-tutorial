package root

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
	controller.Get("/", d.handler.GetAll)
	return controller
}
