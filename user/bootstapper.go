package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"school/controllers"
	"net/http"
)

type status struct {
	Status string `json:"status"`
}

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	 AllowOrigins: []string{"*"},
	 AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.GET("/users/", controllers.HandleGet)
	e.POST("/user", controllers.HandlePost)
	e.GET("/healthcheck", health)

   // Server
   e.Logger.Fatal(e.Start(":1323"))
}

func health(c echo.Context) error {
	health := &status{Status:"UP"}
	return c.JSON(http.StatusOK, health)
}
