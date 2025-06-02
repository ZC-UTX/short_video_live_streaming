package service

import (
	"context"
	__ "github.com/zchengutx/testproject/topics"
	"topic_server/handler/model"
	"topic_server/utlis"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedTopicServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) ListTopic(_ context.Context, in *__.ListTopicReq) (*__.ListTopicResp, error) {

	var topic model.VideoTopic

	video, err := topic.ListVideo()
	if err != nil {
		return nil, err
	}

	var topicList []*__.ListsTopic

	// 初始化算法
	algo := utlis.NewHotTopicAlgorithm()

	// 获取热门话题 (每页10条)
	hotTopics := algo.GetHotTopics(video, 1, 10)

	for _, videoTopic := range hotTopics {

		count := utlis.SumCount(videoTopic.Id)

		listsTopic := __.ListsTopic{
			Title:      videoTopic.Title,
			CreateUser: int64(videoTopic.CreateUser),
			HotScore:   float32(videoTopic.HotScore),
			VideoSum:   float32(count[videoTopic.Id]),
		}

		topicList = append(topicList, &listsTopic)
	}
	return &__.ListTopicResp{
		Code:  200,
		Msg:   "list topic success",
		Lists: topicList,
	}, nil
}
