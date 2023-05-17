package main

import (
	"github.com/eulbyvan/go-auth-jwt/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin router
	r := gin.Default()

	// setup login route
	r.POST("/auth/login", auth.LoginHandler)

	userRouter := r.Group("api/v1/users")
	// middleware
	userRouter.Use(auth.AuthMiddleware())

	// setup get user profile route
	userRouter.GET("/:id/profile", auth.ProfileHandler)

	// start serverr
	r.Run(":8080")
}



