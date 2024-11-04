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
		// adminSignInRouter := authRouter.Group("/admin") //LJIIVNtrSjVO
		// {
		// 	adminSignInRouter.POST("/sign-in", controllers.AuthController.LoginAdmin)
		// }
		authRouter.POST("/sign-in", controllers.AuthController.Signin)
	}

	// router.Use(middleware.JWTAuth(app.Env.AccessTokenSecret, app.Redis))
	// logoutRouter := router.Group("/logout")
	// {
	// 	logoutRouter.POST("", controllers.UserController.Logout)
	// }

	// usersRouter := router.Group("/users")
	// {
	// 	usersRouter.GET("",
	// 		middleware.RBACMiddleware(controllers.LoginController.UserRepository, models.ADMIN, ""),
	// 		controllers.UsersController.GetAll)
	// 	usersRouter.GET("/:userId/profile",
	// 		middleware.RBACMiddleware(controllers.LoginController.UserRepository, models.ADMIN, ""),
	// 		controllers.UsersController.GetProfile)
	// }

	// userRouter := router.Group("/user")
	// {
	// 	userRouter.GET("/profile", controllers.UserController.GetProfile)
	// 	userRouter.PATCH("/profile", controllers.UserController.UpdateProfile)
	// 	userRouter.DELETE("/profile", controllers.UserController.DeleteProfile)
	// }
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
