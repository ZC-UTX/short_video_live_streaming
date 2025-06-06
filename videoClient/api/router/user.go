package router

import (
	"api/handler/api"
	"github.com/gin-gonic/gin"
)

func UserModel(group *gin.RouterGroup) {
	user := group.Group("/user")
	{
		user.POST("register", api.Register)
	}
}
