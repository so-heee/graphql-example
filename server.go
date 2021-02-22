package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kelseyhightower/envconfig"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/so-heee/echo-graphql-example/graph"
	"github.com/so-heee/echo-graphql-example/graph/generated"
)

type Env struct {
	DBUser                 string `required:"true" split_words:"true"`
	DBPass                 string `required:"true" split_words:"true"`
	InstanceConnectionName string `required:"true" split_words:"true"`
}

var env Env

func main() {

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s", env.DBUser, env.DBPass, env.InstanceConnectionName))
	if err != nil {
		log.Fatalln(err)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	// e.GET("/", hello)

	// gqlgen
	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func init() {

	if err := envconfig.Process("", &env); err != nil {
		log.Fatal("fail to load config wiht env :", err)
	}

}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World2!\n")
}