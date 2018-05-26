package wxsrv

import (
	"GraduationDesign/msgtype"
	"bytes"
	"encoding/json"
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
