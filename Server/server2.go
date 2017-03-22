package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var database = "VCRS"
var user = "test"
var password = "test"

type staffs struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func adminLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	pass := r.FormValue("password")
	db, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err = db.Ping(); err != nil {
		log.Print(err)
		return
	}
	var name string
	defer db.Close()
	row := db.QueryRow("SELECT username FROM admin WHERE password=? AND username=?", pass, username)
	e := row.Scan(&name)
	if e != nil {
		log.Println(e)
		return
	}
	if name == username {
		q := "select * from staff;"
		rows, errs := db.Query(q)
		if errs != nil {
			log.Print(errs)
		}
		defer rows.Close()
		list := make([]staffs, 0, 100)
		var name, pass string
		for rows.Next() {
			err := rows.Scan(&name, &pass)
			if err != nil {
				log.Print(err)
			}
			o := &staffs{Username: name, Password: pass}
			list = append(list, *o)
		}
		ljson, je := json.Marshal(list)
		if je != nil {
			log.Print(je)
		}
		fmt.Fprintf(w, string(ljson))
	}
}

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.HandleFunc("/adminLogin/", adminLoginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
`
