package main

import (
	usersControllers "auth_with_token/controllers"
	"net/http"

	"auth_with_token/models"

	auth "auth_with_token/auth"

	"github.com/gin-gonic/gin"

	"auth_with_token/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/session/cookie"
)

func main() {
	models.GetConnect()
	models.Seeder()

	r := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/login", func(ctx *gin.Context) {
		auth.LoginHandler(ctx, models.DB)
		username := ctx.PostForm("username")
		// TODO: Lakukan validasi username dan password jika diperlukan

		// Menyimpan username dalam session
		session := sessions.Default(ctx)
		session.Set("username", username)
		errSess := session.Save()

		if errSess != nil {
			ctx.JSON(http.StatusOK, gin.H{"message": "Login Invalid"})
			ctx.AbortWithStatus(http.StatusBadGateway)
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Login Success"})
		}
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
