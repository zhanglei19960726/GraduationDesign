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
    "button": [
        {
            "type": "view", 
            "name": "课件下载", 
            "url": "http://www.baidu.com"
        }, 
        {
            "name": "数据库教程", 
            "sub_button": [
                {
                    "type": "view", 
                    "name": "mysql教程", 
                    "url": "http://www.runoob.com/mysql/mysql-tutorial.html"
                }, 
                {
                    "type":"view",
					"name":"sql server 教程",
					"url":"http://www.runoob.com/sql/sql-tutorial.html"
                }
            ]
        }
    ]
}`
	accesstoken, err := GetAndUpdateDBWxAToken()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	fmt.Println("token is :", accesstoken)
	return pushWxMenuCreate(accesstoken, []byte(menuStr))
}
