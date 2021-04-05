package main

import (
	"log"

	"github.com/ivanwang123/go-blog/router"
	"github.com/ivanwang123/go-blog/stores"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	store, err := stores.NewStore("postgres://postgres:postgres@localhost/go_blog?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("store", store)
			return next(c)
		}
	})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.Router(e)

	e.Logger.Fatal(e.Start(":8080"))
}
