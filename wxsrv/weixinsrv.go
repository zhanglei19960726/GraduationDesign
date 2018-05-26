package wxsrv

import (
	"fmt"
	"github.com/wizjin/weixin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

var (
	goPath   = os.Getenv("GOPATH")
	htmlPath = goPath + "/src/GraduationDesign/html/"
)

//
const (
	token           = "zhang"
	appID           = "wxf4b1e3a9d5753984"
	appSecret       = "c8981b2fc40b3ecc24f22dc644829099"
	studyOnlineKey  = "Mykey001"
	talkSpace       = "Mykey002"
	sqlKey          = "Mykey003"
	sqlModlekey     = "Mykey004"
	redirectUri     = "http://www.zhangleispace.club/upload"
	sqlPictureURL   = "http://mmbiz.qpic.cn/mmbiz_jpg/gLxmiaSTpZo1dJVGVgic7L2VBqzoFxanCPO9z7HMcqJO3t1tOMHYqbtpgEp1icj3lib6nDj89T4GyRHwo1Dzb881dw/0"
	modlePictureURL = "http://mmbiz.qpic.cn/mmbiz_jpg/gLxmiaSTpZo1dJVGVgic7L2VBqzoFxanCPpb5SFr2sxdD1OletbgblLICK9Hwt8lqZFh57x6IZINsJKicu5rRYYlw/0"
)

//文本消息的处理函数
func echo(w weixin.ResponseWriter, r *weixin.Request) {
	content := ""
	if r.Content == "学习" {
		content = "请输入以下内容获取学习内容：\r\nSQL语言,数据库安全性和完整性，数据库模式"
	} else {
		content = r.Content
	}
	w.ReplyText(content)
}

//关注事件的处理函数
func subscribe(writer weixin.ResponseWriter, request *weixin.Request) {
	writer.ReplyText("欢迎关注")
	wx := writer.GetWeixin()
	articles := make([]weixin.Article, 2)
	articles[0].Title = "整体情况"
	articles[1].Title = "教学大纲"
	writer.PostNews(articles)
	createMenu(wx)
}

//创建菜单
func createMenu(wx *weixin.Weixin) error {
	menu := &weixin.Menu{make([]weixin.MenuButton, 3)}
	menu.Buttons[0].Name = "在线学习"
	menu.Buttons[0].SubButtons = make([]weixin.MenuButton, 2)
	menu.Buttons[0].SubButtons[0].Name = "sql 语句"
	menu.Buttons[0].SubButtons[0].Key = sqlKey
	menu.Buttons[0].SubButtons[0].Type = weixin.MenuButtonTypeKey
	menu.Buttons[0].SubButtons[1].Name = "数据库模式"
	menu.Buttons[0].SubButtons[1].Key = sqlModlekey
	menu.Buttons[0].SubButtons[1].Type = weixin.MenuButtonTypeKey
	menu.Buttons[1].Name = "精彩案例"
	menu.Buttons[1].SubButtons = make([]weixin.MenuButton, 2)
	menu.Buttons[1].SubButtons[0].Name = "mysql教程"
	menu.Buttons[1].SubButtons[0].Type = weixin.MenuButtonTypeUrl
	menu.Buttons[1].SubButtons[0].Url = "http://www.runoob.com/mysql/mysql-tutorial.html"
	menu.Buttons[1].SubButtons[1].Name = "sql server 教程"
	menu.Buttons[1].SubButtons[1].Type = weixin.MenuButtonTypeUrl
	menu.Buttons[1].SubButtons[1].Url = "http://www.runoob.com/sql/sql-tutorial.html"
	menu.Buttons[2].Name = "在线指导"
	menu.Buttons[2].Type = weixin.MenuButtonTypeUrl
	menu.Buttons[2].Url = "http://www.runoob.com/sql/sql-tutorial.html"
	err := wx.CreateMenu(menu)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

//接收点击菜单跳转链接时的事件
func eventView(writer weixin.ResponseWriter, request *weixin.Request) {
	wx := writer.GetWeixin()
	articles := make([]weixin.Article, 3)
	if request.EventKey == sqlKey {
		sqlURL := wx.CreateRedirectURL("http://www.zhangleispace.club/sql", weixin.RedirectURLScopeBasic, "")
		articles[0].Title = "sql 语句"
		articles[0].PicUrl = sqlPictureURL
		articles[0].Description = "zhangleihaha"
		articles[0].Url = sqlURL
		articles[1].Title = "数据库模型"
		articles[1].PicUrl = modlePictureURL
		articles[2].Title = "数据库完整性和安全性"
		articles[2].PicUrl = modlePictureURL
		writer.ReplyNews(articles)
	} else if request.EventKey == talkSpace {
	}
}

func userAgree(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("send msg err : ", err.Error())
		return
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("get msg err: ", err.Error())
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("read body err :", err.Error())
		return
	}
	fmt.Println(string(body))
}

func sqlHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(htmlPath + "sql.html")
	if err != nil {
		log.Println(err.Error())
		return
	}
	t.Execute(w, nil)
}

func Run() {
	mux := weixin.New(token, appID, appSecret)
	//注册文本消息函数
	mux.HandleFunc(weixin.MsgTypeText, echo)
	//注册关注函数
	mux.HandleFunc(weixin.MsgTypeEventSubscribe, subscribe)
	//注册点击事件
	mux.HandleFunc(weixin.MsgTypeEventClick, eventView)
	http.Handle("/", mux)
	http.HandleFunc("/sql", sqlHandler)
	http.ListenAndServe(":80", nil)
}
