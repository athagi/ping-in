package main

import (
	"github.com/athagi/src/ping-in/app"
	"github.com/goadesign/goa"
)

// URIController implements the uri resource.
type URIController struct {
	*goa.Controller
}

// NewURIController creates a uri controller.
func NewURIController(service *goa.Service) *URIController {
	return &URIController{Controller: service.NewController("URIController")}
}

// Host runs the host action.
func (c *URIController) Host(ctx *app.HostURIContext) error {
	// URIController_Host: start_implement

	// Put your logic here

	return nil
	// URIController_Host: end_implement
}
