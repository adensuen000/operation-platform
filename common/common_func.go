package common

import (
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"os"
	"time"
)

// 判断目录或者文件存在
func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// 创建文件
func makeFileExist(fileName string) (bool, error) {
	err := os.MkdirAll(fileName, 0755)
	if err != nil {
		errStr := fmt.Sprintf("创建目录%s失败: ", fileName)
		logger.Error(errors.New(errStr + err.Error()))
		return false, errors.New(errStr + err.Error())
	}
	return true, nil
}

// 获取当前时间并进行时间格式化
func TimeFormat() (time.Time, error) {
	// 加载 UTC 时区
	utcLoc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println("加载 UTC 时区失败:", err)
		return time.Time{}, errors.New(fmt.Sprintf("设置时区失败: ", err.Error()))
	}
	t := time.Now().In(utcLoc)
	timeStr := t.Format("2006-01-02 15:04:05")
	// 将字符串解析为时间
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return time.Time{}, errors.New(fmt.Sprintf("字符串解析为时间失败: ", err))
	}
	return parsedTime, nil
}

// 将数据写进数据库前进行时间字段的转换
func TimeConvert(t time.Time) (time.Time, error) {

	return time.Time{}, nil
}

// 将unix时间戳转换为日期格式
func UnixToDate(ts int64) (time.Time, error) {
	t1 := time.Unix(ts, 0)
	// 加载 UTC 时区
	utcLoc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println("加载 UTC 时区失败:", err)
		return time.Time{}, errors.New(fmt.Sprintf("设置时区失败: ", err.Error()))
	}
	t2 := t1.In(utcLoc)
	timeStr := t2.Format("2006-01-02 15:04:05")
	// 将字符串解析为时间
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return time.Time{}, errors.New(fmt.Sprintf("字符串解析为时间失败: ", err))
	}
	return parsedTime, nil
}
