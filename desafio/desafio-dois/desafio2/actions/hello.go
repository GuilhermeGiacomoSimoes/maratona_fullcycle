package actions

import (

	"net/http"
	"github.com/gobuffalo/bufalo"
)



func HelloFullCycle( c buffalo.Context ) error {
	return c.Render(http.StatusOK, r.HTML( "hello.html" ))
}
