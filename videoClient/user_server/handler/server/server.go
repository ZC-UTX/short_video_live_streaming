package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
	__ "user_server/basic/proto"
	"user_server/handler/dao"
	"user_server/handler/model"
	"user_server/pkg"
)

// server is used to implement helloworld.GreeterServer.
type UserServer struct {
	__.UnimplementedUserServer
}

// // SayHello implements helloworld.GreeterServer
func (s *UserServer) Register(_ context.Context, in *__.RegisterReq) (*__.RegisterResp, error) {
	var user model.VideoUser
	// 根据手机号查找用户
	err := dao.GetOneByFields(&model.VideoUser{Mobile: in.Mobile}, &user)
	if err != false {
		// 用户不存在，创建新用户
		user = model.VideoUser{
			Name:          in.Name,
			NickName:      pkg.RandomString(12),         // 随机昵称
			UserCode:      pkg.RandomFleetingString(12), // 随机用户编码
			Signature:     in.Signature,
			Sex:           in.Sex,
			IpAddress:     in.IpAddress,
			Constellation: in.Constellation,
			Status:        in.Status,
			AvatorFileId:  in.AcatorFileId,
			AuthriryInfo:  in.AuthriryInfo,
			Mobile:        in.Mobile,
			RealNameAuth:  in.RealNameAuth,
			Age:           in.Age,
			OnlineStatus:  in.OnlineStatus,
			AuthrityType:  in.AuthrityType,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}
		if err = dao.Create(&user); err != false {
			return nil, fmt.Errorf("创建用户失败: %v", err)
		}
	} else {
		// 用户已存在，更新用户信息
		user.Name = in.Name
		user.Signature = in.Signature
		user.Sex = in.Sex
		user.IpAddress = in.IpAddress
		user.Constellation = in.Constellation
		user.AvatorFileId = in.AcatorFileId
		user.Age = in.Age
		user.OnlineStatus = in.OnlineStatus
		user.UpdatedAt = time.Now()

		// 如果昵称为空，生成随机昵称
		if user.NickName == "" {
			user.NickName = pkg.RandomString(12)
		}
		if err = dao.Update(&user); err != false {
			return nil, fmt.Errorf("更新用户信息失败: %v", err)
		}
	}
	return &__.RegisterResp{
		Message:  "登录成功",
		Code:     http.StatusOK,
		UserId:   int64(user.Id),
		UserCode: user.UserCode, // 返回用户账号
	}, nil
}
