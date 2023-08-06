package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, th *UserHandler) {
	api := r.Group("/auth")
	{
		api.POST("/signIn", th.LoginUser)
		api.POST("/signUp", th.RegisterUser)
	}
}
