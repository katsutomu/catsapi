// @APIVersion 1.0.0
// @APITitle catAPI
// @APIDescription catAPIだよ

package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	v1 := e.Group("/v1")
	{
		v1.GET("/cats", func(c echo.Context) error {
			t := c.QueryParam("type")
			i, err := strconv.Atoi(t)
			log.Print(err)
			if err != nil {
				return err
			}
			name := SearchCatAction(i)
			return c.String(http.StatusOK, name)
		})
	}
	e.Logger.Fatal(e.Start(":1323"))
}
