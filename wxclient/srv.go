package wxclient

import (
	"html/template"
	"net/http"
	"os"
)

var (
	filePath = "\\src\\GraduationDesign\\html\\"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	goPath := os.Getenv("GOPATH")
	t, err := template.ParseFiles(goPath + filePath + "admin.html")
	if err != nil {
		panic(err)
		return
	}
	t.Execute(w, nil)
	if r.Method == "POST" {
		title := r.FormValue("title")
		author := r.FormValue("author")
		digest := r.FormValue("digest")
		content := r.FormValue("content")
		w.Write([]byte("title is :" + title + " autor is :" + author + " digest is :" + digest + " content:" + content))
	}
}

func HomeHanler(w http.ResponseWriter, r *http.Request) {
	goPath := os.Getenv("GOPATH")
	if r.RequestURI == "/admin.html" {
		http.Redirect(w, r, "/admin", http.StatusFound)
	} else {
		t, err := template.ParseFiles(goPath + filePath + "home.html")
		if err != nil {
			panic(err)
			return
		}
		t.Execute(w, nil)
	}
}
