package middle

import (
	"time"
)

// 把字符串格式的时间参数转换为time.Time格式
func TimeSwitch(tstr string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", tstr)
	return t
}
