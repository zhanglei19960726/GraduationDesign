package diyMen

import (
	"GraduationDesign/client/WxPlatUtil"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//创建菜单微信返回json格式
type MenErrorResponse struct {
	ErrorCode int
	ErrMsg    string
}

var (
	menuFetchUrl = " https://api.weixin.qq.com/cgi-bin/menu/create?access_token="
)

func pushWxMenuCreate(accessToken string, menuJsonBytes []byte) error {
	requsLine := menuFetchUrl + accessToken
	resp, err := http.Post(requsLine, "application/json; encoding=utf-8", bytes.NewReader(menuJsonBytes))
	if err != nil {
		log.Println("发送建立菜单请求失败：", err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("发送post请求创建菜单读取返回body错误", err)
		return err
	}
	atr := &MenErrorResponse{}
	err = json.Unmarshal(body, &atr)
	if err != nil {
		log.Println("发送post请求创建菜单解析body错误", err)
		return err
	}
	fmt.Println("发送Post请求获取 微信返回的错误信息", atr)
	return fmt.Errorf("%s", atr.ErrMsg)
}

func CreateWxMenu() error {
	menuStr := ` {
     "button":[
     {    
          "type":"click",
          "name":"今日歌曲",
          "key":"V1001_TODAY_MUSIC"
      },
      {
           "name":"菜单",
           "sub_button":[
           {    
               "type":"view",
               "name":"搜索",
               "url":"http://www.soso.com/"
            },
            {
                 "type":"miniprogram",
                 "name":"wxa",
                 "url":"http://mp.weixin.qq.com",
                 "appid":"wx286b93c14bbf93aa",
                 "pagepath":"pages/lunar/index"
             },
            {
               "type":"click",
               "name":"赞一下我们",
               "key":"V1001_GOOD"
            }]
       }]
 	}`
	err := WxPlatUtil.GetAndUpdateDBWxAToken()
	fmt.Println("11111111111111111111111111", WxPlatUtil.Accesstoken)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return pushWxMenuCreate(WxPlatUtil.Accesstoken, []byte(menuStr))
}
