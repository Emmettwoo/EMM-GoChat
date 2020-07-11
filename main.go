package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseBody struct {
	Code int `json:"code"`
	Message string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func main() {
	// 登入接口 监听
	http.HandleFunc("/user/login", userLogin)

	// 启动Web服务器
	http.ListenAndServe(":8080", nil)
}

// 登入接口 实现
func userLogin (writer http.ResponseWriter, request *http.Request) {
	// 获取用户名密码
	request.ParseForm()
	username := request.PostForm.Get("username")
	password := request.PostForm.Get("password")

	// 匹配用户名密码
	loginSuccess := false
	if username == "admin" && password == "admin" {
		loginSuccess = true
	}

	if loginSuccess {
		// 登入成功的逻辑处理
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "login success"

		ResponseSend(writer, 0, data, "登入成功")
	} else {
		ResponseSend(writer, -1, nil, "账号或密码错误")
	}
}

// 返回响应数据
func ResponseSend(writer http.ResponseWriter, code int, data interface{}, message string) {
	// 拼接响应数据
	responseBody := ResponseBody{
		Code: code,
		Message:  message,
		Data: data,
	}

	responseJson, err := json.Marshal(responseBody)
	if err != nil {
		log.Println("Response Json Generate Failed")
	}

	// 返回数据到前端
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseJson)
}