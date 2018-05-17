package wxclient

import (
	"GraduationDesign/msgtype"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	addNewsUrl    = "https://api.weixin.qq.com/cgi-bin/material/add_news"
	addPictureUrl = "https://api.weixin.qq.com/cgi-bin/material/add_material"
)

func doPost(accessToken string, newBytes []byte) (*msgtypetype.ArticlesResp, error) {
	postReq, err := http.NewRequest("POST", strings.Join([]string{addNewsUrl, "?access_token=", accessToken}, ""), bytes.NewReader(newBytes))
	if err != nil {
		fmt.Println("向微信新增永久素材建立请求失败", err)
		return nil, err
	}
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
	fmt.Println(token)
	id, err := doPost(token, req)
	return id.MediaId, err
}

func AddPicture(fileName string) error {
	buf := make([]byte, 0)
	file, err := os.OpenFile(goPath+picturePath+fileName, os.O_RDONLY, 0666)
	if err != nil {
		panic(err.Error())
		return err
	}
	defer file.Close()
	num, err := file.Read(buf)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("1111111111111111111", num)
	reqBody := msgtypetype.MaterialReq{
		Media:    buf,
		Filename: fileName,
	}
	data, _ := json.Marshal(reqBody)
	fmt.Println(string(data))
	_, err = GetAndUpdateDBWxAToken()
	if err != nil {
		panic(err.Error())
		return err
	}
	//postReq, err := http.NewRequest("POST", strings.Join([]string{addNewsUrl, "?access_token=", token, "&type=", "image"}, ""), bytes.NewReader(data))
	//client := &http.Client{}
	//resp, err := client.Do(postReq)
	//if err != nil {
	//	fmt.Println("client向微信新增图片素材建立请求失败", err)
	//	return err
	//} else {
	//	fmt.Println("向微信新增图片素材建立成功")
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Println("读取消息失败")
	//	return err
	//}
	//fmt.Println(string(body))
	return nil
}
