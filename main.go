//go:generate goagen bootstrap -d ping-in/design

package main

import (
	"github.com/athagi/src/ping-in/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("adder")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "operands" controller
	c := NewOperandsController(service)
	app.MountOperandsController(service, c)
	// Mount "test" controller
	c2 := NewTestController(service)
	app.MountTestController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
