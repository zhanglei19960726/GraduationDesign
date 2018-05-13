package wxclient

import (
	"GraduationDesign/wxsrv"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func doPost(accessToken string, newBytes []byte) (*ArticlesResp, error) {
	postReq, err := http.NewRequest("POST",
		"https://api.weixin.qq.com/cgi-bin/material/add_news?access_token="+accessToken,
		bytes.NewReader(newBytes))

	if err != nil {
		fmt.Println("向微信新增永久素材建立请求失败", err)
		return nil, err
	}

	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		fmt.Println("client向微信新增永久素材建立请求失败", err)
		return nil, err
	} else {
		fmt.Println("向微信新增永久素材建立成功")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.Status != "200 OK" {
		resperr := &wxsrv.MenErrorResponse{}
		json.Unmarshal(body, resperr)
		return nil, errors.New("error code" + resperr.ErrorCode + " error msge" + resperr.ErrMsg)
	}
	if err != nil {
		log.Println("读取消息失败")
		return nil, err
	}
	media := &ArticlesResp{}
	err = json.Unmarshal(body, media)
	return media, err
}

func AddNews() (string, error) {
	news := &ArticlesReq{}
	news.Title = "zhanglei"
	news.Content = "zhanglei"
	news.ThumbMediaId = ""
	news.ShowCoverPic = 1
	news.ContentSourceUrl = "http://www.baidu.com"
	err := GetAndUpdateDBWxAToken()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	req, _ := json.Marshal(news)
	id, err := doPost(Accesstoken, req)
	return id.MediaId, err
}
