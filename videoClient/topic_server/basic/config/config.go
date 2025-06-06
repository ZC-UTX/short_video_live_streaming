package config

type AppConfig struct {
	Mysql     Mysql
	Redis     Redis
	ALiYun    ALiYun
	BaiDuYun  BaiDuYun
	TenXunYun TenXunYun
}

type Mysql struct {
	User     string `json:"User"`
	Password string `json:"Password"`
	Host     string `json:"Host"`
	Port     int    `json:"Port"`
	Database string `json:"Database"`
}

type Redis struct {
	Addr     string `json:"Addr"`
	Password string `json:"Password"`
	Db       int    `json:"Db"`
}

type ALiYun struct {
	AccessKeyID     string `json:"AccessKeyID"`
	AccessKeySecret string `json:"AccessKeySecret"`
}
type BaiDuYun struct {
	AccessKeyAPI    string `json:"AccessKeyAPI"`
	AccessKeySecret string `json:"AccessKeySecret"`
}

type TenXunYun struct {
	SecretID  string `json:"SecretID"`
	SecretKey string `json:"SecretKey"`
}
