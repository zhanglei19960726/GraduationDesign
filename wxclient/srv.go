package wxclient

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	goPath      = os.Getenv("GOPATH")
	filePath    = "/src/GraduationDesign/html/"
	picturePath = "/src/GraduationDesign/file/"
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
			fmt.Println("hahah")
		}
		f, err := os.OpenFile(goPath+picturePath+handle.Filename, os.O_RDWR|os.O_CREATE, 0777)
		io.Copy(f, file)
		AddPicture(handle.Filename)
		defer f.Close()
		w.Write([]byte("title is :" + title + " autor is :" + author + " digest is :" + digest + " content:" + content + " imagei is:" + handle.Filename))
	}
}

func HomeHanler(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/admin.html" {
		http.Redirect(w, r, "/admin", http.StatusFound)
	} else if r.RequestURI == "/upload.html" {
		http.Redirect(w, r, "/upload", http.StatusFound)
	} else {
		t, err := template.ParseFiles(goPath + filePath + "home.html")
		if err != nil {
			panic(err)
			return
		}
		t.Execute(w, nil)
	}
}

func Run() {
	http.HandleFunc("/", HomeHanler)
	http.HandleFunc("/admin", AdminHandler)
	http.ListenAndServe(":8081", nil)
}

//import (
//	"fmt"
//	"html/template"
//	"io"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"os"
//)
//
//
//)
//
//func AdminHandler(w http.ResponseWriter, r *http.Request) {
//	t, err := template.ParseFiles(goPath + filePath + "admin.html")
//
//	if err != nil {
//		panic(err)
//		return
//	}
//	t.Execute(w, nil)
//	if r.Method == "POST" {
//		title := r.FormValue("title")
//		author := r.FormValue("author")
//		digest := r.FormValue("digest")
//		content := r.FormValue("content")
//		file, handle, err := r.FormFile("image")
//		if err != nil {
//			log.Println(err.Error())
//		}
//		f, err := os.OpenFile(goPath+picturePath+handle.Filename, os.O_RDWR|os.O_CREATE, 0777)
//		io.Copy(f, file)
//		f.Close()
//		err = AddPicture(handle.Filename)
//		if err != nil {
//			log.Println(err.Error())
//		}
//		w.Write([]byte("title is :" + title + " autor is :" + author + " digest is :" + digest + " content:" + content + " imagei is:" + handle.Filename))
//	}
//}
//
//func HomeHanler(w http.ResponseWriter, r *http.Request) {
//	if r.RequestURI == "/admin.html" {
//		http.Redirect(w, r, "/admin", http.StatusFound)
//	} else if r.RequestURI == "/upload.html" {
//		http.Redirect(w, r, "/upload", http.StatusFound)
//	} else {
//		t, err := template.ParseFiles(goPath + filePath + "home.html")
//		if err != nil {
//			panic(err)
//			return
//		}
//		t.Execute(w, nil)
//	}
//}
//
//func GetData(w http.ResponseWriter, r *http.Request) {
//	resp, err := http.Get("http://www.zhangleispace.club:8009/images/")
//	if err != nil {
//		panic(err.Error())
//		return
//	}
//	body, _ := ioutil.ReadAll(resp.Body)
//	defer resp.Body.Close()
//	w.Write(body)
//	fmt.Println(string(body))
//}
