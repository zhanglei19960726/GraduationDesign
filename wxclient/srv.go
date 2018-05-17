package wxclient

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	goPath      = os.Getenv("GOPATH")
	filePath    = "/src/GraduationDesign/html/"
	picturePath = "\\src\\GraduationDesign\\picture\\"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
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
		file, handle, err := r.FormFile("image")
		if err != nil {
			log.Println(err.Error())
		}
		f, err := os.OpenFile(goPath+picturePath+handle.Filename, os.O_RDWR|os.O_CREATE, 0777)
		io.Copy(f, file)
		f.Close()
		err = AddPicture(handle.Filename)
		if err != nil {
			log.Println(err.Error())
		}
		w.Write([]byte("title is :" + title + " autor is :" + author + " digest is :" + digest + " content:" + content + " imagei is:" + handle.Filename))
	}
}

func HomeHanler(w http.ResponseWriter, r *http.Request) {
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
