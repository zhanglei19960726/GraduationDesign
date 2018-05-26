package wxsrv

import (
	"log"
	"net/http"
	"text/template"
)

func sqlHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(htmlPath + "sql.html")
	if err != nil {
		log.Println(err.Error())
		return
	}
	t.Execute(w, nil)
}

func moduleHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(htmlPath + "modle.html")
	if err != nil {
		log.Println(err.Error())
		return
	}
	t.Execute(w, nil)
}

func sqlSerHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(htmlPath + "sqlSer.html")
	if err != nil {
		log.Println(err.Error())
		return
	}
	t.Execute(w, nil)
}
