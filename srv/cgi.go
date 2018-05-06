package srv

import (
	"fmt"
	"net/http"
	"strings"
)

func handle(w http.ResponseWriter, r *http.Request) {
	msgType := strings.Join(r.Form["MsgType"], "")
	switch msgType {
	case "test":
		content := strings.Join(r.Form["Content"], "")
		fmt.Println(content)
	}
}
