package service

import (
	"context"
	__ "works_server/basic/works"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedWorksServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(_ context.Context, in *__.CreatedWorksReq) (*__.CreatedWorksResp, error) {
	return &__.CreatedWorksResp{
		Code: 0,
		Msg:  "",
	}, nil
}
