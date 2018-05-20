package wxsrv

import (
	"fmt"
	"github.com/wizjin/weixin"
	"log"
	"net/http"
	"os"
)

var (
	goPath   = os.Getenv("GOPATH")
	filePath = goPath + "/src/GraduationDesign/file/"
)

const (
	token                   = "zhang"
	appID                   = "wxf4b1e3a9d5753984"
	appSecret               = "c8981b2fc40b3ecc24f22dc644829099"
	databaseIntroductionKey = "Mykey001"
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
	wx := writer.GetWeixin()
	createMenu(wx)
}

//创建菜单
func createMenu(wx *weixin.Weixin) error {
	menu := &weixin.Menu{make([]weixin.MenuButton, 2)}
	menu.Buttons[0].Name = "数据库简介"
	menu.Buttons[0].Type = weixin.MenuButtonTypeKey
	menu.Buttons[0].Key = databaseIntroductionKey
	menu.Buttons[1].Name = "数据库教程"
	menu.Buttons[1].SubButtons = make([]weixin.MenuButton, 2)
	menu.Buttons[1].SubButtons[0].Name = "mysql教程"
	menu.Buttons[1].SubButtons[0].Type = weixin.MenuButtonTypeUrl
	menu.Buttons[1].SubButtons[0].Url = "http://www.runoob.com/mysql/mysql-tutorial.html"
	menu.Buttons[1].SubButtons[1].Name = "sql server 教程"
	menu.Buttons[1].SubButtons[1].Type = weixin.MenuButtonTypeUrl
	menu.Buttons[1].SubButtons[1].Url = "http://www.runoob.com/sql/sql-tutorial.html"
	err := wx.CreateMenu(menu)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

//接收点击菜单跳转链接时的事件
func eventView(writer weixin.ResponseWriter, request *weixin.Request) {
	if request.EventKey == databaseIntroductionKey {
		article := make([]weixin.Article, 1)
		article[0].Title = "test"
		article[0].Description = "数据库(Database)是按照数据结构来组织、存储和管理数据的仓库，" +
			"它产生于距今六十多年前，随着信息技术和市场的发展，特别是二十世纪九十年代以后，" +
			"数据管理不再仅仅是存储和管理数据，而转变成用户所需要的各种数据管理的方式。数据库有很多种类型，" +
			"从最简单的存储有各种数据的表格到能够进行海量数据存储的大型数据库系统都在各个方面得到了广泛的应用。" +
			"在信息化社会，充分有效地管理和利用各类信息资源，是进行科学研究和决策管理的前提条件。" +
			"数据库技术是管理信息系统、办公自动化系统、决策支持系统等各类信息系统的核心部分，" +
			"是进行科学研究和决策管理的重要技术手段"
		article[0].PicUrl = goPath + filePath + "/1.png"
		err := writer.PostNews(article)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func reciveMessage(w weixin.ResponseWriter, r *weixin.Request) (mediaId string, err error) {
	// 上传本地文件并获取MediaID
	mediaId, err = w.UploadMediaFromFile(weixin.MediaTypeImage, filePath+"/1.png")
	return
}

func Run() {
	mux := weixin.New(token, appID, appSecret)
	//注册文本消息函数
	mux.HandleFunc(weixin.MsgTypeText, echo)
	//注册关注函数
	mux.HandleFunc(weixin.MsgTypeEventSubscribe, subscribe)
	//接收点击菜单跳转链接时的事件
	mux.HandleFunc(weixin.MsgTypeEventClick, eventView)
	http.Handle("/", mux)
	http.ListenAndServe(":80", nil)
}
