package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/gRPC-example/gateway/auth"

	_ "github.com/nurzzaat/gRPC-example/gateway/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(router *gin.Engine) {
	controllers, err := SetupControllers()
	if err != nil {
		log.Println("Error handling routes:", err.Error())

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/sign-in", controllers.AuthController.Signin)
		authRouter.POST("/sign-up", controllers.AuthController.Signup)
	}
}

type controllers struct {
	AuthController *auth.AuthController
}

func SetupControllers() (controllers, error) {
	controllers := controllers{}

	authClient, err := auth.NewAuthClient()
	if err != nil {
		return controllers, err
	}

	controllers.AuthController = auth.NewAuthController(authClient)
	return controllers, nil
}
