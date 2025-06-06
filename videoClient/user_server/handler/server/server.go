package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/zchengutx/testproject/SendSms"
	"github.com/zchengutx/testproject/config"
	_ "github.com/zchengutx/testproject/init-redis"
	"github.com/zchengutx/testproject/nickName"
	__ "github.com/zchengutx/testproject/user"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"user_server/handler/model"
)

// server is used to implement helloworld.GreeterServer.
type UserServer struct {
	__.UnimplementedUserServer
}

// 短信验证码
func (s *UserServer) SendSms(_ context.Context, in *__.SendSmsReq) (*__.SendSmsResp, error) {
	code := rand.Intn(9000) + 1000
	sms, err := SendSms.SendSms(in.Mobile, strconv.Itoa(code))
	if err != nil {
		return nil, err
	}
	if *sms.Body.Code != "OK" {
		return nil, errors.New(*sms.Body.Message)
	}
	config.Rdb.Set(context.Background(), "sendSms"+in.Mobile+in.Source, strconv.Itoa(code), time.Minute*5)
	return &__.SendSmsResp{}, nil
}

// 登录注册一体化
// // SayHello implements helloworld.GreeterServer
func (s *UserServer) RegisterOrLogin(ctx context.Context, in *__.RegisterReq) (*__.RegisterResp, error) {
	// 1. 检查短信验证码（Redis）
	smsCode, err := config.Rdb.Get(ctx, "sendSms"+in.Mobile+"register").Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("请先获取短信验证码")
	} else if err != nil {
		return nil, fmt.Errorf("验证码校验失败: %v", err)
	}
	if smsCode != in.SmsCode { // 假设 in.SmsCode 是前端传入的验证码
		return nil, fmt.Errorf("短信验证码错误")
	}

	// 2. 查询用户是否存在
	var user model.VideoUser
	model.GetOneByFields(&model.VideoUser{Mobile: in.Mobile}, &user)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("查询用户失败: %v", err)
	}

	// 3. 用户不存在 → 注册
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = model.VideoUser{
			Mobile:   in.Mobile,
			NickName: nickName.RandomString(12),         // 默认随机昵称
			UserCode: nickName.RandomFleetingString(10), // 用户唯一标识
		}
		if model.Create(&user) {
			return nil, fmt.Errorf("注册失败: %v", err)
		}

		return &__.RegisterResp{
			Message:  "注册成功",
			Code:     http.StatusOK,
			UserId:   int64(user.Id),
			UserCode: user.UserCode,
		}, nil
	}

	// 4. 用户已存在 → 登录
	// 更新最后登录时间（可选）
	user.UpdatedAt = time.Now()
	if user.NickName == "" {
		user.NickName = nickName.RandomString(12)
	}
	if !model.Update(&user) {
		return nil, fmt.Errorf("更新用户信息失败: %v", err)
	}

	return &__.RegisterResp{
		Message:  "登录成功",
		Code:     http.StatusOK,
		UserId:   int64(user.Id),
		UserCode: user.UserCode,
	}, nil
}
