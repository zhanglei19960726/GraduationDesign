package wxsrv

import (
	"html/template"
	"net/http"
)

var templates map[string]*template.Template

//
//func init() {
//	if templates == nil {
//		templates = make(map[string]*template.Template)
//	}
//
//	templates["index"] = template.Must(template.ParseFiles("template/index.html"))
//}

func renderTemplate(w http.ResponseWriter, name string, viewModel interface{}) {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["index"] = template.Must(template.ParseFiles("resource/template/index.html"))
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.Execute(w, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}
