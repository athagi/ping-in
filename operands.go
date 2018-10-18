package main

import (
	"strconv"

	"github.com/athagi/src/ping-in/app"
	"github.com/goadesign/goa"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("OperandsController")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	// OperandsController_Add: start_implement

	// Put your logic here
	sum := ctx.Left + ctx.Right
	return ctx.OK([]byte(strconv.Itoa(sum)))
	// OperandsController_Add: end_implement
}
