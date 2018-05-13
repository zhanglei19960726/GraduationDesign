package main

import (
	"GraduationDesign/wxclient"
	"fmt"
)

func main() {
	wxclient.Run()
	wxclient.GetAndUpdateDBWxAToken()
	fmt.Println(wxclient.Accesstoken)
}
