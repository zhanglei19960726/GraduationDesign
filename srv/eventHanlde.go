package srv

func clickHanlde(eventKey string) string {
	switch eventKey {
	case "V1":
		return "课件"
	case "V2":
		return "安装教程"
	case "V3":
		return "赞一下我们"
	case "V4":
		return "赞一下我们"
	default:
		return "error"
	}
}

func subscribeHandle(eventKey string) string {
	return "hello"
}

func unsubscribeHanlde(eventKey string) {

}
