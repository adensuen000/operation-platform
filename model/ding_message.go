package model

//设置工单消息通知的格式体

type DingMessage struct {
	//MessageID  int64 `json:"message_id"`
	At at `json:"at"`
	//Link       link
	//Markdown   markdown
	Text    text   `json:"text"`
	Msgtype string `json:"msgtype"`
	//CreateTime time.Time `json:"create_time"`
}

// @的人员
type at struct {
	IsAtAll   string   `json:"isAtAll"`
	AtMobiles []string `json:"atMobiles"`
}

// 链接消息
type link struct {
	MessageUrl string `json:"messageUrl"`
	PicUrl     string `json:"picUrl"`
	Text       string `json:"text"`
	Title      string `json:"title"`
}

// markdown消息
type markdown struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

// 文本消息
type text struct {
	Content string `json:"content"`
}
