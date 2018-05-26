package wxclient

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
	id, err := doPost(token, req)
	return id.MediaId, err
}

func AddPicture(fileName string) error {
	token, err := GetAndUpdateDBWxAToken()
	fmt.Println(token)
	if err != nil {
		panic(err.Error())
		return err
	}
	//bodyBuf := &bytes.Buffer{}
	//bodyWriter := multipart.NewWriter(bodyBuf)
	//fileWriter, err := bodyWriter.CreateFormFile("media", fileName)
	//if err != nil {
	//	fmt.Println("error writing to buffer")
	//	return err
	//}
	//file, err := os.Open(goPath + picturePath + fileName)
	//buf := make([]byte, 10240)
	//file.Read(buf)
	//fmt.Println("buf is ", buf)
	//_, err = io.Copy(fileWriter, bytes.NewReader(buf))
	//if err != nil {
	//	panic(err.Error())
	//	return err
	//}
	//fmt.Println(fileWriter)
	//contentType := bodyWriter.FormDataContentType()
	//defer bodyWriter.Close()
	//resp, err := http.Post(strings.Join([]string{"https://api.weixin.qq.com/cgi-bin/media/uploadimg", "?access_token=", token}, ""), contentType, bodyBuf)
	//if err != nil {
	//	log.Println(err.Error())
	//	return err
	//}
	//respBody, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Println(err.Error())
	//	return err
	//}
	//defer resp.Body.Close()
	//fmt.Println(string(respBody))
	return nil
}
