package wxsrv

import (
	"encoding/json"
	"fmt"
	"github.com/wizjin/weixin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	token               = "zhang"
	appID               = "wxf4b1e3a9d5753984"
	appSecret           = "c8981b2fc40b3ecc24f22dc644829099"
	sqlKey              = "Mykey001"
	sqlModlekey         = "Mykey002"
	sqlSerKey           = "Mykey003"
	aboutKey            = "Mykey004"
	modlePicMedia       = "YREDkCL6wmBhl3cwhtjCFKxxlBy8btTVwu7OygZd5YU"
	modleNewsMedia      = "YREDkCL6wmBhl3cwhtjCFANV6ULcTfUZX5ZpOnNS7RM"
	teachAboutPicMedia  = "YREDkCL6wmBhl3cwhtjCFLgghKx774P7iLNHyZ-um84"
	teachAboutNewsMedia = "YREDkCL6wmBhl3cwhtjCFF8_Tj95EmkbOtabTZOGxrs"
	zhengtiPicMedia     = "YREDkCL6wmBhl3cwhtjCFPnJ3605TeQB-HvBMQCa3uM"
	zhengtiNewsMedia    = "YREDkCL6wmBhl3cwhtjCFGDIKsFhr_jzwJbr90HhkY0"
	sqlPicMedia         = "YREDkCL6wmBhl3cwhtjCFJMvm1nzupbiq12IhstEmWg"
	sqlNewsMedia        = "YREDkCL6wmBhl3cwhtjCFEodwmDmkeOqhoxPT3Vgot0"
	sqlSerNewsMia       = "YREDkCL6wmBhl3cwhtjCFL_jY3h4d1lDXkDePVM0KFk"
	zhengtiPicURL       = "http://mmbiz.qpic.cn/mmbiz_jpg/gLxmiaSTpZo0xAiaRg1TLh358xmTibKFMVUawp2ialic2G3qmw8xTDvXEEOwHiakbht2UqqnKiaUXn6LtkObnaGIIjo1Q/0?wx_fmt=jpeg"
	sqlPictureURL       = "http://mmbiz.qpic.cn/mmbiz_jpg/gLxmiaSTpZo1dJVGVgic7L2VBqzoFxanCPO9z7HMcqJO3t1tOMHYqbtpgEp1icj3lib6nDj89T4GyRHwo1Dzb881dw/0?wx_fmt=jpeg"
	teachAboutPicURL    = "http://mmbiz.qpic.cn/mmbiz_jpg/gLxmiaSTpZo0xAiaRg1TLh358xmTibKFMVUjYe0guqWNz3s6fRTo1dUjGIPKXD09OlSRs4v0AOdG40bgh4O5gm1jg/0?wx_fmt=jpeg"
	modlePictureURL     = "http://mmbiz.qpic.cn/mmbiz_jpg/gLxmiaSTpZo1dJVGVgic7L2VBqzoFxanCPpb5SFr2sxdD1OletbgblLICK9Hwt8lqZFh57x6IZINsJKicu5rRYYlw/0?wx_fmt=jpeg"
	sqlSerNewsURL       = "http://mp.weixin.qq.com/s?__biz=MzU5NTU4MTIyMw==&mid=100000020&idx=1&sn=4001746daa6a881e54795172d49287a2&chksm=7e6e815c4919084a58bf927d1cd9ae890653906c2bfd9e17db63d79c313c0c7819955236ab6f#rd"
	teachNewsURL        = "http://mp.weixin.qq.com/s?__biz=MzU5NTU4MTIyMw==&mid=100000022&idx=1&sn=3692c5a939be07dd62c2db732e234273&chksm=7e6e815e49190848e7fe37b25dae2b3b6649435693ed797cda28c053ecd40797a60071f93c13#rd"
	zhengtiNewsURL      = "http://mp.weixin.qq.com/s?__biz=MzU5NTU4MTIyMw==&mid=100000024&idx=1&sn=7e5e64fdd124a2723b5b68859cc84330&chksm=7e6e815049190846f60911e6d14abeb3faf27453531c797a2ac8cadd40297a285515fa970c0c#rd"
	sqlNewsURL          = "http://mp.weixin.qq.com/s?__biz=MzU5NTU4MTIyMw==&mid=100000015&idx=1&sn=ad97809ea27f19c7653ef70b33df9379&chksm=7e6e814749190851f5c0a34d145473e2655ba50038e631ef76b5f32559dcf4cd684ff3d9ca6b#rd"
	modleNewsURL        = "http://mp.weixin.qq.com/s?__biz=MzU5NTU4MTIyMw==&mid=100000019&idx=1&sn=7dcedc05d52750da06db22d54fbb7591&chksm=7e6e815b4919084d81a3e663b98b71f04aec96931ec201ac8f75078b258aaa1c18df3e49727f#rd"
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
		sendOneArticle(w, "数据库安全性和完整性", modlePictureURL, sqlSerNewsURL, "")
	case "数据库模式":
		sendOneArticle(w, "数据库模式", modlePictureURL, modleNewsURL, "")
	default:
		if strings.Contains(r.Content, "音乐+") == true {
			muisc := strings.Split(r.Content, "+")
			info, err := getMuisc(muisc[1])
			if err != nil {
				w.ReplyText(err.Error())
				return
			}
			sendOneArticle(w, info.SongInfo.Title+info.SongInfo.Author, info.SongInfo.PicURL, info.Bitrate.FileLink, "")
		} else {
			content = "回复“学习”，获取学习内容\r\n" +
				"上传地理位置获取天气状况\r\n" +
				"回复“音乐+歌曲名”，获取歌曲"
			w.ReplyText(content)
		}
	}

}

//关注事件的处理函数
func subscribe(writer weixin.ResponseWriter, request *weixin.Request) {
	wx := writer.GetWeixin()
	articles := make([]weixin.Article, 2)
	articles[0].Title = "整体情况"
	articles[0].PicUrl = zhengtiPicURL
	articles[0].Url = zhengtiNewsURL
	articles[1].Title = "教学大纲"
	articles[1].PicUrl = teachAboutPicURL
	articles[1].Url = teachNewsURL
	writer.PostNews(articles)
	createMenu(wx)
}

//创建菜单
func createMenu(wx *weixin.Weixin) error {
	menu := &weixin.Menu{Buttons: make([]weixin.MenuButton, 3)}
	menu.Buttons[0].Name = "学习"
	menu.Buttons[0].SubButtons = make([]weixin.MenuButton, 5)
	menu.Buttons[0].SubButtons[0].Name = "sql 语句"
	menu.Buttons[0].SubButtons[0].Key = sqlKey
	menu.Buttons[0].SubButtons[0].Type = weixin.MenuButtonTypeKey
	menu.Buttons[0].SubButtons[1].Name = "数据库模式"
	menu.Buttons[0].SubButtons[1].Key = sqlModlekey
	menu.Buttons[0].SubButtons[1].Type = weixin.MenuButtonTypeKey
	menu.Buttons[0].SubButtons[2].Name = "数据库安全性和完整性"
	menu.Buttons[0].SubButtons[2].Key = sqlSerKey
	menu.Buttons[0].SubButtons[2].Type = weixin.MenuButtonTypeKey
	menu.Buttons[0].SubButtons[3].Name = "关于我们"
	menu.Buttons[0].SubButtons[3].Type = weixin.MenuButtonTypeKey
	menu.Buttons[0].SubButtons[3].Key = aboutKey
	menu.Buttons[0].SubButtons[4].Name = "在线学习"
	menu.Buttons[0].SubButtons[4].Type = weixin.MenuButtonTypeUrl
	menu.Buttons[0].SubButtons[4].Url = "http://www.zhangleispace.club/upload"
	menu.Buttons[1].Name = "精彩案例"
	menu.Buttons[1].SubButtons = make([]weixin.MenuButton, 2)
	menu.Buttons[1].SubButtons[0].Name = "mysql教程"
	menu.Buttons[1].SubButtons[0].Type = weixin.MenuButtonTypeUrl
	menu.Buttons[1].SubButtons[0].Url = "http://www.runoob.com/mysql/mysql-tutorial.html"
	menu.Buttons[1].SubButtons[1].Name = "sql server 教程"
	menu.Buttons[1].SubButtons[1].Type = weixin.MenuButtonTypeUrl
	menu.Buttons[1].SubButtons[1].Url = "http://www.runoob.com/sql/sql-tutorial.html"
	menu.Buttons[2].Name = "放松一刻"
	menu.Buttons[2].Type = weixin.MenuButtonTypeUrl
	menu.Buttons[2].Url = "http://music.163.com/"
	err := wx.CreateMenu(menu)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

//接收点击菜单跳转链接时的事件
func eventView(writer weixin.ResponseWriter, request *weixin.Request) {
	switch request.EventKey {
	case sqlKey:
		articles := make([]weixin.Article, 1)
		articles[0].Title = "sql 语句"
		articles[0].PicUrl = sqlPictureURL
		articles[0].Url = sqlNewsURL
		writer.ReplyNews(articles)
	case sqlModlekey:
		articles := make([]weixin.Article, 1)
		articles[0].Title = "数据库模式"
		articles[0].PicUrl = modlePictureURL
		articles[0].Url = modleNewsURL
		writer.ReplyNews(articles)
	case sqlSerKey:
		articles := make([]weixin.Article, 1)
		articles[0].Title = "数据库安全性和完整性"
		articles[0].PicUrl = modlePictureURL
		articles[0].Url = sqlSerNewsURL
		writer.ReplyNews(articles)
	case aboutKey:
		articles := make([]weixin.Article, 2)
		articles[0].Title = "整体情况"
		articles[0].PicUrl = zhengtiPicURL
		articles[0].Url = zhengtiNewsURL
		articles[1].Title = "教学大纲"
		articles[1].PicUrl = teachAboutPicURL
		articles[1].Url = teachNewsURL
		writer.ReplyNews(articles)
	default:
		writer.ReplyOK()
	}
}

func location(writer weixin.ResponseWriter, request *weixin.Request) {
	x := strconv.FormatFloat(float64(request.LocationX), 'f', 6, 64)
	y := strconv.FormatFloat(float64(request.LocationY), 'f', 6, 64)
	response, err := http.Get("https://free-api.heweather.com/s6/weather/now?location=" + x + "," + y + "&key=bef3e2e4c99a4884ae76299f5fc9d407")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer response.Body.Close()
	buf, _ := ioutil.ReadAll(response.Body)
	data := &HeWeather6s{}
	err = json.Unmarshal(buf, data)
	if err != nil {
		log.Println(err.Error())
		return
	}
	he := data.HeWeather6[0]
	content := "更新时间：" + he.Update.Loc + "" +
		"\r\n" + he.Basic.AdminArea + he.Basic.ParentCity + he.Basic.Location + "" +
		"\r\n" + "实时天气状况：" + he.Now.ContText + "" +
		"\r\n" + "体感温度：" + he.Now.Fl + "度" +
		"\r\n" + "温度：" + he.Now.Tmp + "度" +
		"\r\n" + "风向：" + he.Now.WindDir + "" +
		"\r\n" + "风力：" + he.Now.WindSc + "" +
		"\r\n" + "风速：" + he.Now.WindSpd + "km/h" +
		"\r\n" + "相对湿度：" + he.Now.Hum + "" +
		"\r\n" + "降水量：" + he.Now.Pcpn + "毫米" +
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
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("resource"))))
	http.HandleFunc("/upload", uploadHandler)
	//http.HandleFunc("/notes/add", addNote)
	//http.HandleFunc("/notes/save", saveNote)
	//http.HandleFunc("/notes/edit/{id}", editNote)
	//http.HandleFunc("/notes/update/{id}", updateNote)
	//article := make([]msgtypetype.Articles, 1)
	//article[0].Title = "整体情况"
	//article[0].ThumbMediaId = zhengtiPicMedia
	//article[0].Content = "整体情况"
	//mediaID, err := AddNews(article)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//fmt.Println(mediaID)
	//fmt.Println(GetAndUpdateDBWxAToken())
	//err := GetNeverExpirePic()
	//if err != nil {
	//	log.Println(err.Error())
	//	return
	//}
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
