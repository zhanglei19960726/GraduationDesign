package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type HomeDB struct {
	CreatTime int64
	Title     string
	Content   string
}

type NaDB struct {
	CreatTime int64
	Title     string
	Content   string
}

type Ke struct {
	CreatTime int64
	Title     string
	Content   string
}

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("mysql", "root:zl961228@tcp(gz-cdb-mvv6udug.sql.tencentcdb.com:62519)/234?charset=utf8")
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
}

func GetHome() (home []HomeDB, err error) {
	rows, err := db.Query("SELECT creattime,title,content FROM  home ")
	if err != nil {
		log.Println(err.Error())
		return
	}
	temp := HomeDB{}
	for rows.Next() {
		if err = rows.Scan(&temp.CreatTime, &temp.Title, &temp.Content); err != nil {
			log.Println(err.Error())
			rows.Close()
			return
		}
		home = append(home, temp)
	}
	rows.Close()
	return
}

func AddHome(creatTime int64, title, content string) error {
	_, err := db.Exec("INSERT  INTO home(creattime,title,content) VALUE (?,?,?)", creatTime, title, content)
	if err != nil {
		panic(err.Error())
	}
	return err
}

func GetKe() (ke []Ke, err error) {
	rows, err := db.Query("SELECT creattime,title,content FROM  ke ")
	if err != nil {
		log.Println(err.Error())
		return
	}
	temp := Ke{}
	for rows.Next() {
		if err = rows.Scan(&temp.CreatTime, &temp.Title, &temp.Content); err != nil {
			log.Println(err.Error())
			rows.Close()
			return
		}
		ke = append(ke, temp)
	}
	rows.Close()
	return
}

func AddKe(creatTime int64, title, content string) error {
	_, err := db.Exec("INSERT  INTO home(creattime,title,content) VALUE (?,?,?)", creatTime, title, content)
	if err != nil {
		panic(err.Error())
	}
	return err
}

func GetNa() (na []NaDB, err error) {
	rows, err := db.Query("SELECT creattime,title,content FROM  na ")
	if err != nil {
		log.Println(err.Error())
		return
	}
	temp := NaDB{}
	for rows.Next() {
		if err = rows.Scan(&temp.CreatTime, &temp.Title, &temp.Content); err != nil {
			log.Println(err.Error())
			rows.Close()
			return
		}
		na = append(na, temp)
	}
	rows.Close()
	return
}

func AddNa(creatTime int64, title, content string) error {
	_, err := db.Exec("INSERT  INTO home(creattime,title,content) VALUE (?,?,?)", creatTime, title, content)
	if err != nil {
		panic(err.Error())
	}
	return err
}
