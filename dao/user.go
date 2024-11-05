package dao

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/wonderivan/logger"
	"operations-platform/db"
	"operations-platform/model"
)

var User user

type user struct {
}

//查询用户

// 新增,用于注册
func (*user) Add(user *model.User) error {
	tx := db.DB.Create(user)
	if tx.Error != nil {
		logger.Error(fmt.Sprintf("添加user失败, %v\n", tx.Error))
		return errors.New(fmt.Sprintf("添加user失败, %v\n", tx.Error))
	}
	return nil
}

// 基于name查询,用于新增
func (*user) Has(name string) (*model.User, bool, error) {
	//初始化要申请内存，不然或报错
	data := &model.User{}
	tx := db.DB.Where("username = ?", name).First(data)
	//如果记录没查到
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, false, nil
	}

	if tx.Error != nil {
		logger.Error(fmt.Sprintf("查询user失败, %v\n", tx.Error))
		return nil, false, errors.New(fmt.Sprintf("查询user失败, %v\n", tx.Error))
	}
	return data, true, nil
}

// 基于token查询,用于中间件token校验
func (*user) GetByToken(token string) (*model.User, bool, error) {
	//初始化要申请内存，不然或报错
	data := &model.User{}
	tx := db.DB.Where("username = ?", token).First(data)
	//如果记录没查到
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, false, nil
	}

	if tx.Error != nil {
		logger.Error(fmt.Sprintf("查询user失败, %v\n", tx.Error))
		return nil, false, errors.New(fmt.Sprintf("查询user失败, %v\n", tx.Error))
	}
	return data, true, nil
}

// 更新token,用于登录
func (*user) UpdateToken(user *model.User, token string) error {
	tx := db.DB.Model(user).Where("username = ? and password = ?", user.Username, user.Password).
		Update("token", token)
	fmt.Println("tx.ERROR: ", tx.Error)

	if tx.Error != nil {
		logger.Error(fmt.Sprintf("更新user token失败, %v\n", tx.Error))
		return errors.New(fmt.Sprintf("更新user token失败, %v\n", tx.Error))
	}
	return nil
}

// 删除用户
func (*user) Delete(name string) (bool, error) {
	// 校验用户是否存在
	res, err := User.Verify(name)
	if err != nil {
		return false, err
	}
	if !res {
		return false, nil
	}
	return true, nil
}

// 校验用户是否存在
func (*user) Verify(name string) (bool, error) {
	_, has, err := User.Has(name)
	if err != nil {
		return false, err
	}
	if !has {
		return false, nil
	}
	return true, nil
}

// 获取用户ID
func (*user) GetUserID(username string) (int, error) {
	singleUser := &model.User{}
	// 校验用户是否存在
	res, err := User.Verify(username)
	if err != nil {
		return 0, err
	}
	if !res {
		return 0, nil
	}

	resQuery := db.DB.Where(" username = ? ", username).First(singleUser)
	if resQuery.Error != nil {
		return 0, errors.New(fmt.Sprintf("查询数据库失败: ", resQuery.Error))
	}
	fmt.Println("singleUser.UserID: ", singleUser.UserID)
	return singleUser.UserID, nil
}

// 根据字符串查询用户名
func (*user) GetUsers(str string) ([]*model.User, error) {
	var (
		users          []*model.User
		queryCondition = "%" + str + "%"
	)
	res := db.DB.Where("username like ? ", queryCondition).Find(&users)
	if res.Error != nil {
		return nil, errors.New(fmt.Sprintf("获取用户列表失败: ", res.Error))
	}
	return users, nil
}
