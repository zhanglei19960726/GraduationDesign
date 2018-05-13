package wxclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type ArticlesReq struct {
	Title            string //标题
	ThumbMediaId     string //图文消息的封面图片素材ID
	Author           string //作者
	Digest           string //图文消息摘要
	ShowCoverPic     int    //是否显示封面
	Content          string //图文消息的具体内容
	ContentSourceUrl string //图文消息的原文地址，即点击“阅读原文”后的URL
}

type ArticlesResp struct {
	MediaId string
}

var (
	addNewsUrl = "https://api.weixin.qq.com/cgi-bin/material/add_news"
)

func doPost(accessToken string, newBytes []byte) (*ArticlesResp, error) {
	postReq, err := http.NewRequest("POST",
		strings.Join([]string{addNewsUrl, "?access_token=", accessToken}, ""),
		bytes.NewReader(newBytes))

	if err != nil {
		fmt.Println("向微信发送菜单建立请求失败", err)
		return nil, err
	}

	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		fmt.Println("client向微信发送菜单建立请求失败", err)
		return nil, err
	} else {
		fmt.Println("向微信发送菜单建立成功")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	media := &ArticlesResp{}
	err = json.Unmarshal(body, media)
	return media, err
}

func AddNews() (string, error) {
	news := `{
"articles": [{
"title": "TITLE",
"thumb_media_id": "THUMB_MEDIA_ID"",
"author": "AUTHOR",
"digest": "DIGEST",
"show_cover_pic": 1,
"content": "CONTENT",
"content_source_url": "http://www.baidu.com"
},
//若新增的是多图文素材，则此处应还有几段articles结构
]
}`
	err := GetAndUpdateDBWxAToken()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	req, _ := json.Marshal(news)
	id, err := doPost(Accesstoken, req)
	return id.MediaId, err
}
