package wxclient

import (
	"io/ioutil"
	"net/http"
	"os"
)

func readHTML() ([]byte, error) {
	var buf []byte
	goPath := os.Getenv("GOPATH")
	file, err := os.Open(goPath + "/src/GraduationDesign/html/admin.html")
	if err != nil {
		return buf, err
	}
	buf, err = ioutil.ReadAll(file)
	if err != nil {
		return buf, err
	}
	return buf, nil
}

func AdminHanler(w http.ResponseWriter, r *http.Request) {
	buf, err := readHTML()
	if err != nil {
		panic(err.Error())
		return
	}
	w.Write(buf)
	if r.Method == "POST" {
		title := r.FormValue("title")
		author := r.FormValue("author")
		digest := r.FormValue("digest")
		content := r.FormValue("content")
		w.Write([]byte("title is :" + title + " autor is :" + author + " digest is :" + digest + " content:" + content))
		AddNews()
	}
}
