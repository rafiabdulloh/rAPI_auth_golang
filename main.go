package main

import (
	usersControllers "auth_with_token/controllers"
	"net/http"

	"auth_with_token/models"

	auth "auth_with_token/auth"

	"github.com/gin-gonic/gin"

	"auth_with_token/middleware"

	// "github.com/gorilla/securecookie"
	
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("very-secret"), []byte("a-lot-secret"))

func main() {
	models.GetConnect()
	models.Seeder()

	r := gin.Default()
	// store := sessions.NewCookieStore([]byte("secret"))
	// Hash keys should be at least 32 bytes long
	// Block keys should be 16 bytes (AES-128) or 32 bytes (AES-256) long.
	// Shorter keys may weaken the encryption used.
	r.Use(func(c *gin.Context) {
        session, _ := store.Get(c.Request, "mysession")
        c.Set("session", session)
        c.Next()
    })

	r.POST("/login", func(ctx *gin.Context) {
		auth.LoginHandler(ctx, models.DB)
		username := ctx.PostForm("username")
		// TODO: Perform username and password validation if needed

		// Storing username in the session
		session := ctx.MustGet("session").(*sessions.Session)
		session.Values["username"] = username
		errSess := session.Save(ctx.Request, ctx.Writer)

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
