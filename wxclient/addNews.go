package wxclient

import (
	"GraduationDesign/msgtype"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func doPost(accessToken string, newBytes []byte) (*msgtypetype.ArticlesResp, error) {
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
	if err != nil {
		log.Println("读取消息失败")
		return nil, err
	}
	fmt.Println("1111111111111", string(body))
	fmt.Println("333333333333333333", bytes.NewReader(newBytes))
	media := &msgtypetype.ArticlesResp{}
	err = json.Unmarshal(body, media)
	return media, err
}

func AddNews() (string, error) {
	news := &msgtypetype.ArticlesReq{}
	news.Title = "zhanglei"
	news.Content = "zhanglei"
	news.ThumbMediaId = "333333333333"
	news.ShowCoverPic = 1
	news.ContentSourceUrl = "http://www.baidu.com"
	err := GetAndUpdateDBWxAToken()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	req, _ := json.Marshal(news)
	fmt.Println(222222222222, string(req))
	id, err := doPost(Accesstoken, req)
	return id.MediaId, err
}
