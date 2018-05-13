package wxclient

import (
	"net/http"
	"os"
)

func readHTML() ([]byte, error) {
	var buf []byte
	file, err := os.Open("../html/admin.html")
	if err != nil {
		return buf, err
	}
	_, err = file.Read(buf)
	if err != nil {
		return buf, err
	}
	return buf, nil
}

func AdminHanler(w http.ResponseWriter, r *http.Request) {

}
