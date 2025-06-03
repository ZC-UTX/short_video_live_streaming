package config

type AppConfig struct {
	Mysql  Mysql
	Redis  Redis
	ALiYun ALiYun
}

type Mysql struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}
type Redis struct {
	Addr     string
	Password string
	DB       int
}
type ALiYun struct {
	AccessKeyID     string
	AccessKeySecret string
}
