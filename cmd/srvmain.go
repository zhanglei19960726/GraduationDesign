package main

import (
	"GraduationDesign/client/diyMen"
	"GraduationDesign/srv"
)

func main() {
	go srv.Run()
	go diyMen.CreateWxMenu()
	select {}
}
