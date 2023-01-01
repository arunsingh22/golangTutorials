package main

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
)

type user struct {
	Name string `json:"name"`
	ID   int    `json:"ID"`
}

//In-Memory Database
var users = map[int]*user{}
var seq = 0

func main() {
	e := echo.New()

	e.POST("/user/:name", func(ctx echo.Context) error {
		u := &user{
			Name: ctx.Param("name"),
			ID:   seq + 1,
		}
		if err := ctx.Bind(u); err != nil {
			return err
		}
		users[u.ID] = u
		seq++
		return ctx.JSON(http.StatusCreated, users)
	})

	// GETs the User with specific ID
	e.GET("/getuser/:id", func(ctx echo.Context) error {
		id, _ := strconv.Atoi(ctx.Param("id"))
		return ctx.JSON(http.StatusOK, users[id])

	})

	// GET All users in Database
	e.GET("/users", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, users)
	})

	// Delete the User with specific ID from Database
	e.DELETE("/delete/:id", func(ctx echo.Context) error {
		id, _ := strconv.Atoi(ctx.Param("id"))
		delete(users, id)
		return ctx.JSON(http.StatusOK, "User Deleted!")
	})

	e.Start(":8000")

}
