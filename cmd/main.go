package main

import (
	"flag"
	"fmt"

	repository "github.com/ThailanTec/rascunho/internal/adapter/repository/database"
	"github.com/ThailanTec/rascunho/internal/adapter/repository/handler"
	"github.com/ThailanTec/rascunho/internal/core/services"
	"github.com/gin-gonic/gin"
)

var (
	repo = flag.String("db", "postgres", "Database for storing messages")
	// redisHost   = "localhost:6379"
	httpHandler *handler.HTTPHandler
	svc         *services.UserService
)

func main() {

	flag.Parse()

	fmt.Printf("Application running using %s\n", *repo)

	switch *repo {
	case "redis":
		/* store := repository.NewMessengerPostgresRepository(redisHost)
		svc = services.NewUserService(store) */
	default:
		store := repository.NewMessengerPostgresRepository()
		svc = services.NewUserService(store)
	}

	InitRoutes()

}

func InitRoutes() {
	router := gin.Default()
	handler := handler.NewHttpHandler(*svc)
	router.GET("/getUser/:name", handler.GetUser)
	router.POST("/create", handler.CreateUser)

	router.Run(":9000")
}
