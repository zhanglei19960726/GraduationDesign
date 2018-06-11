package wxsrv

import (
	"GraduationDesign/db"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/arstd/weixin"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["index"] = template.Must(template.ParseFiles("resource/template/index.html"))
	templates["first"] = template.Must(template.ParseFiles("resource/template/first.html"))
	templates["ho"] = template.Must(template.ParseFiles("resource/template/navigation.html"))
	templates["na"] = template.Must(template.ParseFiles("resource/template/uploadtonggao.html"))
	templates["getHo"] = template.Must(template.ParseFiles("resource/template/getHo.html"))
	templates["getNa"] = template.Must(template.ParseFiles("resource/template/getNa.html"))
	templates["getKe"] = template.Must(template.ParseFiles("resource/template/getKe.html"))
	templates["Ke"] = template.Must(template.ParseFiles("resource/template/Addkejian.html"))
}

func renderTemplate(w http.ResponseWriter, name string, viewModel interface{}) {
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

func reHandler(w http.ResponseWriter, r *http.Request) {
	info, err := getUserInfo()
	if err != nil {
		log.Println(err.Error())
		return
	}
	renderTemplate(w, "first", info)
}

var (
	Note string
)

func hoHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "ho", Note)
}

func SendTemplateMsg(msg *weixin.TemplateMsg) error {
	token, err := GetAndUpdateDBWxAToken()
	if err != nil {
		return err
	}
	data, _ := json.Marshal(msg)
	postReq, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="+token, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取消息失败")
		return err
	}
	return nil
}

func submit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.PostFormValue("title")
	desc := r.PostFormValue("des")
	Note = "提交成功"
	http.Redirect(w, r, "/ho", 302)
	Note = ""
	list, err := getUserList()
	if err != nil {
		log.Println(err.Error())
		return
	}
	for _, v := range list.Data.Openid {
		msg := &weixin.TemplateMsg{
			ToUser:     v,
			TemplateId: "ucQWmyKD2xd6FULnqmiBqYdbeR-xTNMBfyw4CSOSJTQ",
			Data: weixin.TemplateData{
				Keyword1: weixin.KeywordPair{
					Value: title,
				},
				Keyword2: weixin.KeywordPair{
					Value: desc,
				},
			},
		}
		err := SendTemplateMsg(msg)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
	err = db.AddHome(time.Now().Unix(), title, desc)
	if err != nil {
		log.Println(err.Error())
	}
}

func naHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "na", Note)
}

func naSubmitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.PostFormValue("title")
	desc := r.PostFormValue("des")
	Note = "提交成功"
	http.Redirect(w, r, "/na", 302)
	list, err := getUserList()
	if err != nil {
		log.Println(err.Error())
		return
	}
	for _, v := range list.Data.Openid {
		msg := &weixin.TemplateMsg{
			ToUser:     v,
			TemplateId: "IZZwdJ8MJoFd4Tw9FXEV6WZHU3smdAq2pZkDtjRt9uM",
			Data: weixin.TemplateData{
				Keyword1: weixin.KeywordPair{
					Value: title,
				},
				Keyword2: weixin.KeywordPair{
					Value: desc,
				},
			},
		}
		err := SendTemplateMsg(msg)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
	db.AddNa(time.Now().Unix(), title, desc)
}

func getHoHandler(w http.ResponseWriter, r *http.Request) {
	home, _ := db.GetHome()
	renderTemplate(w, "getHo", home)
}

func getNaHandler(w http.ResponseWriter, r *http.Request) {
	na, _ := db.GetNa()
	renderTemplate(w, "getNa", na)
}

func getKeHandler(w http.ResponseWriter, r *http.Request) {
	ke, _ := db.GetKe()
	renderTemplate(w, "getKe", ke)
}

func keHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "ke", nil)
}

func keSubmit(w http.ResponseWriter, r *http.Request) {
	file, head, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//创建文件
	fW, err := os.Create(head.Filename)
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		fmt.Println("文件保存失败")
		return
	}
	//io.WriteString(w, head.Filename+" 保存成功")
	http.Redirect(w, r, "/ke", http.StatusFound)
}
