package actions

import (

	"net/http"
	"github.com/gobuffalo/buffalo"
)



func HelloFullCycle( c buffalo.Context ) error {
	return c.Render(http.StatusOK, r.HTML( "hello.html" ))
}
