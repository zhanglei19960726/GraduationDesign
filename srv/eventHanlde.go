package srv

import "fmt"

func clickHanlde(eventKey string) string {
	fmt.Println("1111111111111", eventKey)
	switch eventKey {
	case "v1":
		return "v1"
	case "v2":
		return "v2"
	case "v3":
		return "v3"
	case "v4":
		return "v4"
	default:
		return "error"
	}
}

func subscribeHandle(eventKey string) string {
	return "hello"
}

func unsubscribeHanlde(eventKey string) {

}
