package diyMen

import (
	"GraduationDesign/client/WxPlatUtil"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//创建菜单微信返回json格式
type MenErrorResponse struct {
	ErrorCode int
	ErrMsg    string
}

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
	menuStr := ` {
    "button": [
        	{
            	"type": "click", 
            	"name": "今日歌曲", 
            	"key": "V1001_TODAY_MUSIC"
        	}, 
        	{
            	"name": "菜单", 
            	"sub_button": [
                	{
                    	"type": "click", 
                    	"name": "赞一下我们", 
                    	"key": "V1001_GOOD"
                	}, 
                	{
                    	"type": "click", 
                    	"name": "赞一下我们", 
                    	"key": "V1001_GOOD"
                	}
            	]
        	},
 			{
            	"name": "菜单", 
				"sub_button": [
                	{
                    	"type": "click", 
                    	"name": "赞一下我们", 
                    	"key": "V1001_GOOD"
                	}, 
                	{
                    	"type": "click", 
                    	"name": "赞一下我们", 
                    	"key": "V1001_GOOD"
                	}
            	]
        	}
    	]
	}`
	err := WxPlatUtil.GetAndUpdateDBWxAToken()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	fmt.Println("111", WxPlatUtil.Accesstoken)
	return pushWxMenuCreate(WxPlatUtil.Accesstoken, []byte(menuStr))
}
