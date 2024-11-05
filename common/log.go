package common

var Log logging

type logging struct {
}

var (
// 日志文件
// LogFile = config.LogPath + "/" + config.LogFile
// LogFile = config.LogFile
)

// 创建日志文件,文件存在则不创建
//func CreateLogFile(name string) (bool, error) {
//	res, err := makeFileExist(LogFile)
//	return res, err
//}

// 记录日志
//func (l logging) Logging(data interface{}) {
//	if _, err := CreateLogFile(LogFile); err != nil {
//		fmt.Println(err)
//	}
//	file, err := os.OpenFile(LogFile, os.O_WRONLY|os.O_APPEND, 755)
//	if err != nil {
//		fmt.Println("打开日志文件失败: ", err.Error())
//	}
//	log.SetOutput(file)
//	log.Println(data)
//}
