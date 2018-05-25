package wxsrv

import (
	"fmt"
	"github.com/wizjin/weixin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	goPath   = os.Getenv("GOPATH")
	filePath = goPath + "/src/GraduationDesign/file/"
)

const (
	token          = "zhang"
	appID          = "wxf4b1e3a9d5753984"
	appSecret      = "c8981b2fc40b3ecc24f22dc644829099"
	studyOnlineKey = "Mykey001"
	talkSpace      = "Mykey002"
	sqlKey         = "Mykey003"
	sqlModlekey    = "Mykey004"
	redirectUri    = "http://www.zhangleispace.club/upload"
)

//文本消息的处理函数
func echo(w weixin.ResponseWriter, r *weixin.Request) {

}

//关注事件的处理函数
func subscribe(writer weixin.ResponseWriter, request *weixin.Request) {
	writer.ReplyText("欢迎关注")
	wx := writer.GetWeixin()
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
	articles := make([]weixin.Article, 1)
	if request.EventKey == sqlKey {
		articles[0].Title = "sql"
		mediaId, err := writer.UploadMediaFromFile(weixin.MediaTypeImage, filePath+"1.png")
		if err != nil {
			log.Println(err.Error())
			return
		}
		articles[0].PicUrl = mediaId
		articles[0].Description = "hahahahaahahhahahhaha"
		err = writer.PostNews(articles)
		if err != nil {
			log.Println(err.Error())
			return
		}
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

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
	if r.Method == "GET" {
		fmt.Println("GET")
	}
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
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":80", nil)
}
