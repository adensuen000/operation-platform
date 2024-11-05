package service

import (
	"errors"
	"github.com/wonderivan/logger"
	"operations-platform/config"
)

var Login login

type login struct {
}

func (l *login) Auth(username, password string) (err error) {
	if username == config.AdminUser && password == config.AdminPwd {
		return nil
	} else {
		logger.Error("登录失败，用户名或者密码错误.")
		return errors.New("登录失败，用户名或者密码错误.")
	}
}
