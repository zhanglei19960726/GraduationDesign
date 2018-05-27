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
	"strings"
)

const (
	addNewsUrl    = "https://api.weixin.qq.com/cgi-bin/material/add_news"
	addPictureUrl = "https://api.weixin.qq.com/cgi-bin/material/add_material"
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
	respone, err := http.Get("http://tingapi.ting.baidu.com/v1/restserver/ting?method=baidu.ting.song.play&songid=8" + id)
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
