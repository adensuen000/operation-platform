package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 创建 HTTP 请求体
	requestBody := []byte(`{"key": "value"}`)

	// 创建 HTTP 请求对象
	req, err := http.NewRequest("POST", "https://example.com/api/endpoint", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("创建 HTTP 请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建自定义的 http.Client
	client := &http.Client{}

	// 发送 HTTPS 请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送 HTTPS 请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		return
	}

	// 打印响应内容
	fmt.Println("响应内容:", string(body))
}
