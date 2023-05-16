package main

import (
	usersControllers "auth_with_token/controllers"

	"auth_with_token/models"

	auth "auth_with_token/auth"

	"github.com/gin-gonic/gin"

	"auth_with_token/middleware"
)

func main() {
	models.GetConnect()
	models.Seeder()

	r := gin.Default()

	r.POST("/login", func(ctx *gin.Context) {
		auth.LoginHandler(ctx, models.DB)
	})

	r.GET("/user", middleware.AuthMiddleware, func(ctx *gin.Context) {
		usersControllers.GetAll(ctx, models.DB)
	})

	r.POST("/user", func(ctx *gin.Context) {
		usersControllers.Post(ctx, models.DB)
	})

	r.GET("/user/:id", func(ctx *gin.Context) {
		usersControllers.GetById(ctx, models.DB)
	})

	r.PUT("/user/:id", func(ctx *gin.Context) {
		usersControllers.Update(ctx, models.DB)
	})
	r.DELETE("/user/:id", func(ctx *gin.Context) {
		usersControllers.Delete(ctx, models.DB)
	})

	// ============================ Running Server =================================

	r.Run()

}