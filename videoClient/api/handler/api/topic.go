package api

import (
	"api/basic/config"
	"github.com/gin-gonic/gin"
	__ "github.com/zchengutx/testproject/topics"

	"net/http"
)

func ListTopic(context *gin.Context) {

	r, _ := config.TopicClient.ListTopic(context, &__.ListTopicReq{})

	if r.Code != 200 {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": r.Code,
			"msg":  r.Msg,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": r.Code,
		"msg":  r.Msg,
		"date": r.Lists,
	})
}
