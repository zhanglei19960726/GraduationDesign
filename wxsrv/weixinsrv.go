package wxsrv

import (
	"encoding/json"
	"fmt"
	"github.com/wizjin/weixin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	token           = "zhang"
	appID           = "wxf4b1e3a9d5753984"
	appSecret       = "c8981b2fc40b3ecc24f22dc644829099"
	sqlKey          = "Mykey001"
	sqlModlekey     = "Mykey002"
	sqlSerKey       = "Mykey003"
	redirectUri     = "http://www.zhangleispace.club/upload"
	modlePicMedia   = "YREDkCL6wmBhl3cwhtjCFKxxlBy8btTVwu7OygZd5YU"
	modleNewsMedia  = "YREDkCL6wmBhl3cwhtjCFJYlD7MKdRyP_1mFoDSBnwY"
	sqlMedia        = "YREDkCL6wmBhl3cwhtjCFJMvm1nzupbiq12IhstEmWg"
	sqlNewsMedia    = "YREDkCL6wmBhl3cwhtjCFEodwmDmkeOqhoxPT3Vgot0"
	sqlNewsURL      = "http://mp.weixin.qq.com/s?__biz=MzU5NTU4MTIyMw==&mid=100000015&idx=1&sn=ad97809ea27f19c7653ef70b33df9379&chksm=7e6e814749190851f5c0a34d145473e2655ba50038e631ef76b5f32559dcf4cd684ff3d9ca6b#rd"
	sqlPictureURL   = "http://mmbiz.qpic.cn/mmbiz_jpg/gLxmiaSTpZo1dJVGVgic7L2VBqzoFxanCPO9z7HMcqJO3t1tOMHYqbtpgEp1icj3lib6nDj89T4GyRHwo1Dzb881dw/0?wx_fmt=jpeg"
	modlePictureURL = "http://mmbiz.qpic.cn/mmbiz_jpg/gLxmiaSTpZo1dJVGVgic7L2VBqzoFxanCPpb5SFr2sxdD1OletbgblLICK9Hwt8lqZFh57x6IZINsJKicu5rRYYlw/0?wx_fmt=jpeg"
	modleNewsURL    = "http://mp.weixin.qq.com/s?__biz=MzU5NTU4MTIyMw==&mid=100000015&idx=1&sn=ad97809ea27f19c7653ef70b33df9379&chksm=7e6e814749190851f5c0a34d145473e2655ba50038e631ef76b5f32559dcf4cd684ff3d9ca6b#rd"
)

func sendOneArticle(w weixin.ResponseWriter, title, picUrl, articleurl, description string) {
	article := make([]weixin.Article, 1)
	article[0].Title = title
	article[0].Description = description
	article[0].PicUrl = picUrl
	article[0].Url = articleurl
	w.ReplyNews(article)
}

//文本消息的处理函数
func echo(w weixin.ResponseWriter, r *weixin.Request) {
	content := ""
	switch r.Content {
	case "学习":
		content = "请输入以下内容获取学习内容：\r\nSQL语言\r\n数据库安全性和完整性\r\n数据库模式"
		w.ReplyText(content)
	case "SQL语言":
		sendOneArticle(w, "SQL 语言", sqlPictureURL, sqlNewsURL, "")
	case "数据库安全性和完整性":
		sendOneArticle(w, "数据库安全性和完整性", modlePictureURL, modleNewsURL, "")
	case "数据库模式":
		sendOneArticle(w, "数据库模式", modlePictureURL, modleNewsURL, "")
	default:
		content = r.Content
		w.ReplyText(content)
	}

}

//关注事件的处理函数
func subscribe(writer weixin.ResponseWriter, request *weixin.Request) {
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
	menu.Buttons[0].SubButtons = make([]weixin.MenuButton, 3)
	menu.Buttons[0].SubButtons[0].Name = "sql 语句"
	menu.Buttons[0].SubButtons[0].Key = sqlKey
	menu.Buttons[0].SubButtons[0].Type = weixin.MenuButtonTypeKey
	menu.Buttons[0].SubButtons[1].Name = "数据库模式"
	menu.Buttons[0].SubButtons[1].Key = sqlModlekey
	menu.Buttons[0].SubButtons[1].Type = weixin.MenuButtonTypeKey
	menu.Buttons[0].SubButtons[2].Name = "数据库安全性和完整性"
	menu.Buttons[0].SubButtons[2].Key = sqlSerKey
	menu.Buttons[0].SubButtons[2].Type = weixin.MenuButtonTypeKey
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
	switch request.EventKey {
	case sqlKey:
		articles[0].Title = "sql 语句"
		articles[0].PicUrl = sqlPictureURL
		articles[0].Url = sqlNewsURL
	case sqlModlekey:
		articles[0].Title = "数据库模式"
		articles[0].PicUrl = modlePictureURL
		articles[0].Url = modleNewsURL
	case sqlSerKey:
		articles[0].Title = "数据库安全性和完整性"
		articles[0].PicUrl = modlePictureURL
		articles[0].Url = modleNewsURL
	}
	writer.ReplyNews(articles)
}

func location(writer weixin.ResponseWriter, request *weixin.Request) {
	fmt.Println(request.LocationX, request.LocationY)
	x := strconv.FormatFloat(float64(request.LocationX), 'f', 6, 64)
	y := strconv.FormatFloat(float64(request.LocationY), 'f', 6, 64)
	fmt.Println(x, y)
	response, err := http.Get("https://free-api.heweather.com/s6/weather/now?location=" + x + "," + y + "&key=bef3e2e4c99a4884ae76299f5fc9d407")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer response.Body.Close()
	buf, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(buf))
	data := &HeWeather6s{}
	err = json.Unmarshal(buf, data)
	if err != nil {
		log.Println(err.Error())
		return
	}
	he := data.HeWeather6[0]
	content := he.Update.Loc + "" +
		"\r\n" + he.Basic.AdminArea + he.Basic.ParentCity + he.Basic.ParentCity + "" +
		"\r\n" + "体感温度：" + he.Now.Fl + "度" +
		"\r\n" + "温度：" + he.Now.Tmp + "度" +
		"\r\n" + "风向：" + he.Now.WindDir + "" +
		"\r\n" + "风力：" + he.Now.WindSc + "" +
		"\r\n" + "风速：" + he.Now.WindSpd + "km/h" +
		"\r\n" + "相对湿度：" + he.Now.Hum + "" +
		"\r\n" + "降水量：" + he.Now.Pcpn + "" +
		"\r\n" + "能见度：" + he.Now.Vis + "公里"
	writer.ReplyText(content)
}

func Run() {
	mux := weixin.New(token, appID, appSecret)
	//注册文本消息函数
	mux.HandleFunc(weixin.MsgTypeText, echo)
	//注册关注函数
	mux.HandleFunc(weixin.MsgTypeEventSubscribe, subscribe)
	//注册点击事件
	mux.HandleFunc(weixin.MsgTypeEventClick, eventView)
	//注册上报地理位置
	mux.HandleFunc(weixin.MsgTypeLocation, location)
	http.Handle("/", mux)
	//article := make([]msgtypetype.Articles, 1)
	//article[0].Title = "数据库模式"
	//article[0].ThumbMediaId = modlePicMedia
	//article[0].Content = "数据库模式"
	//mediaID, err := AddNews(article)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//fmt.Println(mediaID)
	//fmt.Println(GetAndUpdateDBWxAToken())
	http.ListenAndServe(":80", nil)
}
