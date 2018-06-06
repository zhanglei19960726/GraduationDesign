package wxsrv

import (
	"GraduationDesign/msgtype"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	addNewsUrl    = "https://api.weixin.qq.com/cgi-bin/material/add_news"
	addPictureUrl = "https://api.weixin.qq.com/cgi-bin/material/add_material"
	loginExpire   = 123
)

type Basic struct {
	Location   string `json:"location"`    //地区／城市名称
	ParentCity string `json:"parent_city"` //该地区/城市的上级城市
	AdminArea  string `json:"admin_area"`  //该地区/城市所属行政区域
}

type Updata struct {
	Loc string `json:"loc"` //当地时间
}

type Now struct {
	Fl       string `json:"fl"`       //体感温 摄氏度
	Tmp      string `json:"tmp"`      //温度 摄氏度
	WindDir  string `json:"wind_dir"` //风向
	WindSc   string `json:"wind_sc"`  //风力
	WindSpd  string `json:"wind_spd"` //风速 km/h
	Hum      string `json:"hum"`      //相对湿度
	Pcpn     string `json:"pcpn"`     //降水量
	Vis      string `json:"vis"`      //能见度 km
	ContText string `json:"cond_txt"` //实况天气状况
}

type HeWeather6 struct {
	Basic  Basic  `json:"basic"`  //基础信息
	Update Updata `json:"update"` //接口更新时间
	Now    Now    `json:"now"`    //实况天气
}

//天气的
type HeWeather6s struct {
	HeWeather6 []HeWeather6 `json:"HeWeather6"`
}

//音乐的

type Error struct {
	Errno        int32  `json:"errno"`
	ErrorCode    int32  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type MuiscFind struct {
	Error
	Song []Song `json:"song"`
}

type Song struct {
	SongID string `json:"songid"`
}

type Info struct {
	Error
	SongInfo SongInfo `json:"songinfo"`
	Bitrate  Bitrate  `json:"bitrate"`
}

type Bitrate struct {
	FileLink string `json:"file_link"`
}

type SongInfo struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	PicURL string `json:"pic_small"`
}

func findMuisc(music string) (string, error) {
	repose, err := http.Get("http://tingapi.ting.baidu.com/v1/restserver/ting?method=baidu.ting.search.catalogSug&query=" + music)
	if err != nil {
		return "", err
	}
	buf, err := ioutil.ReadAll(repose.Body)
	defer repose.Body.Close()
	muiscFind := &MuiscFind{}
	err = json.Unmarshal(buf, muiscFind)
	if err != nil {
		return "", err
	}
	if muiscFind.ErrorMessage != "" {
		return "", errors.New(muiscFind.ErrorMessage)
	}
	return muiscFind.Song[0].SongID, nil
}

func getMuisc(music string) (info *Info, err error) {
	//搜索歌曲
	id, err := findMuisc(music)
	if err != nil {
		return info, err
	}
	respone, err := http.Get("http://tingapi.ting.baidu.com/v1/restserver/ting?method=baidu.ting.song.play&songid=" + id)
	if err != nil {
		return info, err
	}
	buf, err := ioutil.ReadAll(respone.Body)
	defer respone.Body.Close()
	info = &Info{}
	err = json.Unmarshal(buf, info)
	if err != nil {
		return info, err
	}
	if info.ErrorMessage != "" {
		return info, errors.New(info.ErrorMessage)
	}
	return info, nil
}

func doPost(accessToken, url string, newBytes []byte) (*msgtypetype.ArticlesResp, error) {
	postReq, err := http.NewRequest("POST", strings.Join([]string{url, "?access_token=", accessToken}, ""), bytes.NewReader(newBytes))
	if err != nil {
		fmt.Println("向微信服务器建立请求失败", err.Error())
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		fmt.Println("向微信服务器发送请求失败", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取消息失败")
		return nil, err
	}
	fmt.Println(string(body))
	media := &msgtypetype.ArticlesResp{}
	err = json.Unmarshal(body, media)
	return media, err
}

func AddNews(articles []msgtypetype.Articles) (string, error) {
	requestBody := msgtypetype.ArticlesReq{
		Articles: articles,
	}
	token, err := GetAndUpdateDBWxAToken()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	req, _ := json.Marshal(requestBody)
	id, err := doPost(token, addNewsUrl, req)
	return id.MediaId, err
}

//二维码
type Expire struct {
	ActionName string     `json:"action_name"`
	ActionInfo ActionInfo `json:"action_info"`
}

type ActionInfo struct {
	SceneId int32 `json:"scene_id"`
}

type ExpireResponse struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds string `json:"expire_seconds"`
	Url           string `json:"url"`
}

func getNeverExpire() (ex *ExpireResponse, err error) {
	expire := Expire{
		ActionName: "QR_LIMIT_SCENE",
		ActionInfo: ActionInfo{
			SceneId: loginExpire,
		},
	}
	data, err := json.Marshal(expire)
	if err != nil {
		return ex, err
	}
	token, err := GetAndUpdateDBWxAToken()
	if err != nil {
		return ex, err
	}
	postReq, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token="+token, bytes.NewBuffer(data))
	if err != nil {
		return ex, err
	}
	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		return ex, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取消息失败")
		return ex, err
	}
	fmt.Println(string(body))
	ex = &ExpireResponse{}
	err = json.Unmarshal(body, ex)
	return ex, nil
}

//获取永久二维码
func GetNeverExpirePic(filePath string) error {
	ex, err := getNeverExpire()
	if err != nil {
		return err
	}
	ex.Ticket, err = url.PathUnescape(ex.Ticket)
	if err != nil {
		return err
	}
	resp, err := http.Get("https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + ex.Ticket)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	_, err = file.Write(body)
	if err != nil {
		return err
	}
	return nil
}

//修改永久图文素材
type Articles struct {
	Titile           string `json:"titile"`
	ThumbMediaId     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	ShowCoverPic     int32  `json:"show_cover_pic"`
	Content          string `json:"content"`
	ContentSourceUrl string `json:"content_source_url"`
}
type Media struct {
	MediaID  string   `json:"media_id"`
	Index    string   `json:"index"`
	Articles Articles `json:"articles"`
}

type ErrMsg struct {
	Errcode string `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func SetMedia() {
	media := Media{
		MediaID: zhengtiNewsMedia,
		Articles: Articles{
			Titile: "软件推荐",
		},
	}
	token, err := GetAndUpdateDBWxAToken()
	if err != nil {
		log.Println("get token error:", err.Error())
		return
	}
	data, err := json.Marshal(media)
	if err != nil {
		log.Println("json marshal error:", err.Error())
		return
	}
	resq, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/material/update_news?access_token="+token, bytes.NewBuffer(data))
	if err != nil {
		log.Println("http post error:", err.Error())
		return
	}
	client := &http.Client{}
	resp, err := client.Do(resq)
	if err != nil {
		log.Println("client Do error:", err.Error())
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println("read body error:", err.Error())
		return
	}
	fmt.Println(string(body))
}
