package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	// http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "GET" {
	// 		fmt.Fprintln(w, "Hello World")
	// 	} else if r.Method == "POST" {
	// 		fmt.Fprintln(w, http.StatusOK)
	// 	}
	// })
	// http.HandleFunc("/login/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "The Request Body: %v", r)
	// })
	// http.ListenAndServe("localhost:8000", nil)

	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello To FIRST Page")
	})

	//fake DB
	ProductList := map[string]string{
		"1": "Fridge",
		"2": "Washing-machine",
		"3": "Dumming",
	}
	e.GET("/products", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, ProductList)
	})

	e.GET("/products/:id", func(ctx echo.Context) error {
		for key, val := range ProductList {
			if ctx.Param("id") == key {
				return ctx.JSON(http.StatusOK, val)
			}
		}
		return ctx.JSON(http.StatusBadRequest, "Producet Not Found! ")
	})

	e.POST("/products", func(ctx echo.Context) error {
		type body struct {
			Name string `json:product_name`
		}
		var reqBody body
		if err := ctx.Bind(&reqBody); err != nil {
			return err
		}

		prod := map[string]string{}
		prod[strconv.Itoa((len(ProductList) + 1))] = reqBody.Name
		//appending to the map
		ProductList[strconv.Itoa((len(ProductList) + 1))] = reqBody.Name
		return ctx.JSON(http.StatusOK, prod)

	})

	e.Start(":8080")

}
