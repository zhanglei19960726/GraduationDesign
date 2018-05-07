package diyMen

import (
	"GraduationDesign/client/WxPlatUtil"
	"bytes"
	"log"
	"net/http"
	"strings"
)

func pushWxMenuCreate(accessToken string, menuJsonBytes []byte) error {
	postReq, err := http.NewRequest("POST",
		strings.Join([]string{"https://api.weixin.qq.com/cgi-bin/menu/create", "?access_token=", accessToken}, ""),
		bytes.NewReader(menuJsonBytes))

	if err != nil {
		log.Println("向微信发送菜单建立请求失败", err)
		return err
	}

	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		log.Println("client向微信发送菜单建立请求失败", err)
		return err
	} else {
		log.Println("向微信发送菜单建立成功")
	}
	defer resp.Body.Close()

	return nil
}

func CreateWxMenu() {
	menuStr := `{
            "button": [
            {
                "name": "进入商城",
                "type": "view",
                "url": "http://www.baidu.com/"
            },
            {

                "name":"管理中心",
                 "sub_button":[
                        {
                        "name": "用户中心",
                        "type": "click",
                        "key": "molan_user_center"
                        },
                        {
                        "name": "公告",
                        "type": "click",
                        "key": "molan_institution"
                        }]
            },
            {
                "name": "资料修改",
                "type": "view",
                "url": "http://www.baidu.com/user_view"
            }
            ]
        }`
	accessToken, err := WxPlatUtil.GetAndUpdateDBWxAToken()
	if err != nil {
		panic(err)
	}
	pushWxMenuCreate(accessToken, []byte(menuStr))
}
