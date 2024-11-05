package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/wonderivan/logger"
	"operations-platform/common"
	"operations-platform/dao"
	"operations-platform/model"
)

var User user

type user struct {
}

// 登录
func (*user) Login(name, password string) error {
	//先校验用户是否存在
	data, has, err := dao.User.Has(name)
	if err != nil {
		return err
	}
	if !has {
		return errors.New("该用户不存在.")
	}
	//校验密码
	if data.Password != password {
		return errors.New("用户名或者密码错误.")
	}
	return nil
}

// 更新token
func (*user) UpdateToken(user *model.User) error {
	token := uuid.New().String()
	return dao.User.UpdateToken(user, token)
}

// 新增
func (*user) Add(user *model.User) error {
	//先校验用户是否存在
	if res := User.Verify(user.Username); res {
		logger.Error(fmt.Sprintf("该用户已存在，请重新创建."))
		return errors.New("该用户已存在，请重新创建.")
	}
	//设置时间字段
	t, err := common.TimeFormat()
	if err != nil {
		return err
	}
	user.CreateTime, user.UpdatedTime = t, t

	if errAdd := dao.User.Add(user); errAdd != nil {
		return errAdd
	}
	return nil
}

// 删除用户
func (*user) Delete(name string) (bool, error) {
	res, err := dao.User.Delete(name)
	if err != nil {
		logger.Error(fmt.Sprintf("删除用户失败: ", err.Error()))
		return false, errors.New("删除用户失败: " + err.Error())
	}
	if !res {
		logger.Error(fmt.Sprintf("删除用户失败: 用户不存在"))
		return false, nil
	}
	return true, nil
}

// 修改
func (*user) Update() {

}

// 查询
func (u user) UserQuery(name string) (*model.User, bool, error) {
	data, res, err := dao.User.Has(name)
	if err != nil {
		return nil, false, nil
	}
	if !res {
		return nil, false, nil
	}
	return data, true, nil
}

// 校验用户是否存在
func (*user) Verify(name string) bool {
	//先校验用户是否存在
	_, has, err := dao.User.Has(name)
	if err != nil {
		return false
	}
	if !has {
		return false
	}
	return true
}

// 获取用户ID
func (*user) GetUserID(username string) (int, error) {
	fmt.Println("username: ", username)
	userID, err := dao.User.GetUserID(username)
	if err != nil {
		return userID, err
	}
	return userID, nil
}

// 根据字符串查询用户名
func (*user) GetUsers(str string) ([]*model.User, error) {
	users, err := dao.User.GetUsers(str)
	if err != nil {
		return nil, err
	}
	return users, nil
}
