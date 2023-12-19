package routers

import "github.com/RandySteven/go-e-commerce.git/interfaces"

type Handlers struct {
	interfaces.UserHandler
}

func NewHandlers() *Handlers {
	return nil
}
