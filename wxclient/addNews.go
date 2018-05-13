package wxclient

import (
	"GraduationDesign/msgtype"
	"bytes"
	"encoding/json"
	"errors"
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
	if resp.Status != "200 OK" {
		resperr := &msgtypetype.MenErrorResponse{}
		json.Unmarshal(body, resperr)
		fmt.Println(resperr)
		return nil, errors.New("error code" + resperr.ErrorCode + " error msge" + resperr.ErrMsg)
	}
	fmt.Println("1111111111111", string(body))
	media := &msgtypetype.ArticlesResp{}
	err = json.Unmarshal(body, media)
	return media, err
}

func AddNews() (string, error) {
	news := &msgtypetype.ArticlesReq{}
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
	fmt.Println(222222222222, string(req))
	id, err := doPost(Accesstoken, req)
	return id.MediaId, err
}
