package router

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	v1 := r.Group("v1")
	{
		apiGroup := v1.Group("api")
		{
			TopicModel(apiGroup)
			UserModel(apiGroup)
		}
	}
}
