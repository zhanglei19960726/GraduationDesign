package main

import (
	"github.com/wizjin/weixin"
	"net/http"
)

var (
	token     = "zhang"
	appID     = "wxf4b1e3a9d5753984"
	appSecret = "c8981b2fc40b3ecc24f22dc644829099"
)

//文本消息的处理函数
func echo(w weixin.ResponseWriter, r *weixin.Request) {
	txt := r.Content //获取用户发送的消息
	repaly := ""
	if txt == "1" {
		repaly = "http://www.zhangleispace.club:8009/images/"
	} else if txt == "2" {
		repaly = "15029236434"
	} else {
		repaly = "回复“1”获得课件下载地址\n回复“2”获得联系方式"
	}
	w.ReplyText(repaly)
}

//关注事件的处理函数
func subscribe(writer weixin.ResponseWriter, request *weixin.Request) {
	writer.ReplyText("欢迎关注")
}

func main() {
	mux := weixin.New(token, appID, appSecret)
	//注册文本消息函数
	mux.HandleFunc(weixin.MsgTypeText, echo)
	//注册关注函数
	mux.HandleFunc(weixin.MsgTypeEventSubscribe, subscribe)
	http.Handle("/", mux)
	http.ListenAndServe(":80", nil)
}
