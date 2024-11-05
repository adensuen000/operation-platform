package config

import "time"

const (
	//服务器启动监听端口
	ListAddr = "0.0.0.0:9091"

	//jwt token的加密因子
	JWT_SECRET = "whatsgood"

	//默认账号密码
	AdminUser = "admin"
	AdminPwd  = "123456"
	//数据库配置
	DBType     = "mysql"
	DBHost     = "43.143.93.53"
	DBPort     = 6000
	DBName     = "platform"
	DBUser     = "root"
	DBPassword = "1qaz@WSX"
	LogMode    = true
	//数据库连接池配置
	MaxIdleConns = 10               //最大空闲连接
	MaxOpenConns = 100              //最大连接数
	MaxLifeTime  = 10 * time.Second //最大存活时间

	//日志文件设置
	//LogPath = "./"
	//LogFile = "platform.log"
)
