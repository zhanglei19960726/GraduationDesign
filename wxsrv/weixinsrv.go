package wxsrv

import (
	"fmt"
	"github.com/wizjin/weixin"
	"log"
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

//获取菜单
func DeleteMenu(wx *weixin.Weixin) {
	menu, err := wx.GetMenu()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(menu)
	}
}

//创建菜单
func createMenu(wx *weixin.Weixin) error {
	menu := &weixin.Menu{make([]weixin.MenuButton, 2)}
	menu.Buttons[0].Name = "数据库简介"
	menu.Buttons[0].Type = weixin.MenuButtonTypeKey
	menu.Buttons[0].Key = "Mykey001"
	//menu.Buttons[1].Name = "数据库教程"
	//menu.Buttons[1].SubButtons = make([]weixin.MenuButton, 2)
	//menu.Buttons[1].SubButtons[0].Name = "mysql教程"
	//menu.Buttons[1].SubButtons[0].Type = weixin.MenuButtonTypeUrl
	//menu.Buttons[1].SubButtons[0].Url = "http://www.runoob.com/mysql/mysql-tutorial.html"
	//menu.Buttons[1].SubButtons[1].Name = "sql server 教程"
	//menu.Buttons[1].SubButtons[1].Type = weixin.MenuButtonTypeUrl
	//menu.Buttons[1].SubButtons[1].Url = "http://www.runoob.com/sql/sql-tutorial.html"
	err := wx.CreateMenu(menu)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}
func Run() {
	mux := weixin.New(token, appID, appSecret)
	//注册文本消息函数
	mux.HandleFunc(weixin.MsgTypeText, echo)
	//注册关注函数
	mux.HandleFunc(weixin.MsgTypeEventSubscribe, subscribe)
	wx := &weixin.Weixin{}
	fmt.Println("11111111111111111")
	err := createMenu(wx)
	fmt.Println("hahaha")
	if err != nil {
		log.Println(err.Error())
		return
	}
	DeleteMenu(wx)
	http.Handle("/", mux)
	http.ListenAndServe(":80", nil)
}
