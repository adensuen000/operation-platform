package service

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"net/http"
	"operations-platform/db"
	"operations-platform/model"
	"strconv"
)

var DingMsg dingMsg

type dingMsg struct {
}

var (
	reqHeader  = "Content-Type: application/json"
	token      = "181a6e5f9c434717c0f765ecc510aaf9789ea142f1d6d57e30438458a048224e"
	reqAddress = "https://oapi.dingtalk.com/robot/send?access_token=" + token
)

// 发送https请求
func (*dingMsg) sendHTTPSRequest(message *model.DingMessage) (*http.Response, error) {
	fmt.Println("message: ", message)

	//把消息结构体格式化成json
	jsonData, err := json.Marshal(message)
	if err != nil {
		logger.Error(fmt.Sprintf("消息格式化失败: ", err.Error()))
		return nil, errors.New(fmt.Sprintf("消息格式化失败: ", err.Error()))
	}

	fmt.Println("jsonData: ", string(jsonData))

	// 创建 HTTP 请求对象
	req, _ := http.NewRequest("POST", reqAddress, bytes.NewBuffer(jsonData))

	//设置请求内容
	req.Header.Set("Content-Type", "application/json")

	// 创建自定义的 http.Client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// 发送 HTTPS 请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送 HTTPS 请求失败: %v", err)
	}
	//fmt.Println("resp: ", resp)
	//body, _ := ioutil.ReadAll(resp.Body)
	//var prettyJSON bytes.Buffer
	//err = json.Indent(&prettyJSON, body, "", "\t")
	//fmt.Println("响应内容:")
	//fmt.Println(string(prettyJSON.Bytes()))

	defer resp.Body.Close()
	return resp, nil
}

// 组装钉钉信息
func (*dingMsg) TranToDingMsg(woa *model.WorkOrderAssign) (*model.DingMessage, error) {
	var (
		message       *model.DingMessage
		phoneNumber   string
		contentPrefix = "work_order: 您有新的工单产生，请记得跟进，工单id: "
		content       = contentPrefix + strconv.FormatInt(woa.TicketID, 10)
		msgType       = "text"
	)
	//初始化结构体指针
	message = new(model.DingMessage)

	//获取要通知的用户的手机号
	phoneNumber, err := DingMsg.GetPhoneNumber(woa.RecieveUserID)
	if err != nil {
		return nil, err
	}
	//赋值
	message.At.IsAtAll = "false"
	message.At.AtMobiles = append(message.At.AtMobiles, phoneNumber)
	message.Text.Content = content
	message.Msgtype = msgType

	return message, nil
}

// 根据用户ID获取用户手机号
func (*dingMsg) GetPhoneNumber(userid int) (string, error) {
	var userIns model.User
	res := db.DB.Where("user_id like ? ", userid).Find(&userIns)
	if res.Error != nil {
		return "", errors.New(fmt.Sprintf("获取用户手机号失败: ", res.Error))
	}

	return userIns.PhoneNumber, nil
}

// 发送工单消息
func (*dingMsg) SendWOMsg(woa *model.WorkOrderAssign) (*http.Response, error) {

	DMsg, err := DingMsg.TranToDingMsg(woa)
	if err != nil {
		return nil, err
	}

	resp, err := DingMsg.sendHTTPSRequest(DMsg)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
