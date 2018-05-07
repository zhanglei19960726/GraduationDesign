package main

import (
	"GraduationDesign/WxPlatUtil"
	"fmt"
)

func main() {
	WxPlatUtil.GetAndUpdateDBWxAToken()
	fmt.Println(WxPlatUtil.AccessToken)
}
