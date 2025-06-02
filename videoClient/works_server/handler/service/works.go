package service

import (
	"context"
	__ "github.com/zchengutx/testproject/works"
	"net/http"
	"works_server/basic/config"
	"works_server/handler/model"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedWorksServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) CreateWorks(_ context.Context, in *__.CreatedWorksReq) (*__.CreatedWorksResp, error) {
	works := model.VideoWorks{
		Title:     in.Title,
		Desc:      in.Desc,
		MusicId:   int32(in.MusicId),
		WorkType:  in.WorkType,
		IpAddress: in.IpAddress,
	}

	err := works.CreatedWorks()
	if err != nil {
		return &__.CreatedWorksResp{
			Code: http.StatusInternalServerError,
			Msg:  "create works failed",
		}, nil
	}
	return &__.CreatedWorksResp{
		Code: http.StatusOK,
		Msg:  "create works success",
	}, nil
}
func (s *Server) ListWorks(_ context.Context, in *__.ListWorksReq) (*__.ListWorksResp, error) {
	var works model.VideoWorks
	listWorks, err := works.ListWorks(in.Page, in.PageSize, in.Title, in.WorkType, in.WorkPermission)
	if err != nil {
		return &__.ListWorksResp{
			Code: http.StatusInternalServerError,
			Msg:  "find works failed",
		}, nil
	}
	var list []*__.List
	for _, work := range listWorks {
		list = append(list, &__.List{
			Title:     work.Title,
			Desc:      work.Desc,
			MusicId:   int64(work.MusicId),
			WorkType:  work.WorkType,
			IpAddress: work.IpAddress,
		})
	}
	var count int64
	config.DB.Model(&model.VideoWorks{}).Count(&count)
	return &__.ListWorksResp{
		Code:  http.StatusOK,
		Msg:   "find works success",
		Lists: list,
		Count: count,
	}, nil
}
