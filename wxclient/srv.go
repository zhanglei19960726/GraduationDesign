package wxclient

import (
	"fmt"
	"log"
)

func Run() {
	id, err := AddNews()
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(id)
}
