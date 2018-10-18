package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("adder", func() {
	Title("The adder API")
	Description("A teaser for goa")
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("operands", func() {
	Action("add", func() {
		Routing(GET("add/:left/:right"))
		Description("add returns the sum of the left and right parameters in the response body")
		Params(func() {
			Param("left", Integer, "Left operand")
			Param("right", Integer, "Right operand")
		})
		Response(OK, "text/plain")
	})

})

var _ = Resource("test", func() {
	Action("add", func() {
		Routing(GET("test/"))
		Description("test function")
		Params(func() {
		})
		Response(OK, "text/plain")
	})

})

var _ = Resource("uri", func() {
	Action("host", func() {
		Routing(GET("uri/host/:host_name"))
		Description("add uri to check endpoint")
		Params(func() {
			Param("host_name", String, "hostname")
		})
		Response(OK, "text/plain")
	})
})
