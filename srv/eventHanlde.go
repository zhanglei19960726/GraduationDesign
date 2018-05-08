package srv

func clickHanlde(eventKey string) string {
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

func subscribeHandle(eventKey string) {

}

func unsubscribeHanlde(eventKey string) {

}
