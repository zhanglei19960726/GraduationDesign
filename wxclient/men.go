package wxclient

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func pushWxMenuCreate(accessToken string, menuJsonBytes []byte) error {
	postReq, err := http.NewRequest("POST",
		strings.Join([]string{"https://api.weixin.qq.com/cgi-bin/menu/create", "?access_token=", accessToken}, ""),
		bytes.NewReader(menuJsonBytes))

	if err != nil {
		fmt.Println("向微信发送菜单建立请求失败", err)
		return err
	}

	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		fmt.Println("client向微信发送菜单建立请求失败", err)
		return err
	} else {
		fmt.Println("向微信发送菜单建立成功")
	}
	defer resp.Body.Close()

	return nil
}

func CreateWxMenu() error {
	menuStr := `{
     "button":[
     {    
          "type":"view",
          "name":"百度一下 ",
          "url":"http://www.baidu.com"
      },
      ]
 }`
	err := GetAndUpdateDBWxAToken()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return pushWxMenuCreate(Accesstoken, []byte(menuStr))
}