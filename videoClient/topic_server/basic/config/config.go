package config

type AppConfig struct {
	Mysql Mysql
}

type Mysql struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}
