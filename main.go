package main

import (
	"fmt"
	"github.com/labstack/echo/v4"

	"net/http"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	e := echo.New()
	e.GET("/", HealthCheck)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save_form", saveForm)
	e.POST("/users", usersJson)
	e.Logger.Fatal(e.Start(":1323"))
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	fmt.Println(c.QueryParams())
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// form value
func saveForm(c echo.Context) error {
	fmt.Println(c.FormParams())
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func usersJson(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	//json_map := make(map[string]interface{})
	//
	//err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(json_map)
	return c.JSON(http.StatusCreated, u)
}
