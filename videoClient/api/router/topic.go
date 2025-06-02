package router

import (
	"api/handler/api"
	"github.com/gin-gonic/gin"
)

func TopicModel(group *gin.RouterGroup) {
	topic := group.Group("/topic")
	{
		topic.GET("listTopic", api.ListTopic)
	}
}
