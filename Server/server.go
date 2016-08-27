package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)
var database = "VCRS"
var user = "test"
var password = "test"

type question struct {
	Id string `json:"id"`
	Class string `json:"class"`
	Ques string `json:"question"`
	Option1 string `json:"option1"`
	Option2 string `json:"option2"`
	Option3 string `json:"option3"`
	Option4 string `json:"option4"`
	Ans string `json:"ans"`
}
func isLoginValid(username string, pass string) bool {
	db, err := sql.Open("mysql", user + ":" + password + "@/" + database)
	if err = db.Ping(); err != nil {
		log.Fatal(err);
		return false;
	}
	var name string
	defer db.Close()
	row := db.QueryRow("SELECT username FROM staff WHERE password=? AND username=?", pass, username)
	e := row.Scan(&name)
	if e != nil {
		log.Println(e)
		return false;
	}
	if name == username {
		return true
	}
	return false
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("username")
	pass := r.FormValue("password")
	isValid := isLoginValid(user, pass)
	if isValid {
		fmt.Fprintf(w, "Logged In")	
	}else {
		fmt.Fprintf(w, "INVALID")
	}
}
func getQuesHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	db, err := sql.Open("mysql", user + ":" + password + "@/" + database)
	if err = db.Ping(); err != nil {
		log.Fatal(err);
	}
	var class, ques, op1, op2, op3, op4, ans string
	defer db.Close()
	row := db.QueryRow("SELECT * FROM sampleTest1 WHERE id = ?",id)
	e := row.Scan(&id, &class, &ques, &op1, &op2, &op3, &op4, &ans)
	if e != nil {
		log.Println(e)
	}
	qobj := &question{Id: id, Class: class, Ques: ques, Option1: op1, Option2: op2, Option3: op3, Option4: op4, Ans: ans}
	qjson, je := json.Marshal(qobj)
	if je != nil{
		log.Fatal(je)
	}
	fmt.Fprintf( w, string(qjson)) 
}
func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.HandleFunc("/login/",loginHandler)
	http.HandleFunc("/getQuestions/",getQuesHandler)
	http.ListenAndServe(":80", nil)
}
