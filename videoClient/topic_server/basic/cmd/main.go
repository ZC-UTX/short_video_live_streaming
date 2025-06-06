package main

import (
	"encoding/json"
	"flag"
	"fmt"
	init_mysql "github.com/zchengutx/testproject/init-mysql"
	init_redis "github.com/zchengutx/testproject/init-redis"
	"github.com/zchengutx/testproject/nacos"
	__ "github.com/zchengutx/testproject/topics"
	"google.golang.org/grpc"
	"log"
	"net"
	"topic_server/basic/config"
	"topic_server/handler/service"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	//初始化nacos
	var err error

	config.Nacos, err = nacos.InitNacos("../config.yaml")
	if err != nil {
		config.Log.Fatal(err.Error())
		return
	}

	var conf config.AppConfig

	json.Unmarshal([]byte(config.Nacos), &conf)

	mysql := init_mysql.NewMysql(conf.Mysql.User, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Database)

	config.DB = mysql.InitMysql()

	redis := init_redis.NewRedis(conf.Redis.Addr, conf.Redis.Password, conf.Redis.Db)

	var errs string
	config.RDB, errs = redis.ExampleNewClient()

	if errs != "" {
		config.Log.Fatal(errs)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	__.RegisterTopicServer(s, &service.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
