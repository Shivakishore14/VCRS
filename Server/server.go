package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var database = "VCRS"
var user = "test"
var password = "test"
var db *sql.DB

type question struct {
	ID      string `json:"id"`
	Class   string `json:"class"`
	Ques    string `json:"question"`
	Option1 string `json:"option1"`
	Option2 string `json:"option2"`
	Option3 string `json:"option3"`
	Option4 string `json:"option4"`
	Ans     string `json:"ans"`
	Type    string `json:"type"`
	Level   string `json:"level"`
}
type testList struct {
	Name   string `json:"name"`
	No     string `json:"no"`
	Level  string `json:"level"`
	Status string `json:"status"`
}
type ansLevel struct {
	Ans   string `json:"ans"`
	Level string `json:"level"`
}
type stuList struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type level struct {
	Number int `json:"number"`
	Val    int `json:"val"`
}
type levelSetModel struct {
	Levels   []level `json:"levels"`
	TestName string  `json:"testName"`
}

/*type testListArray struct {
	arr TestList `json:"arr"`
}*/
type response struct {
	TestName string   `json:"testname"`
	Sid      string   `json:"sid"`
	Ans      []string `json:"ans"`
}
type result struct {
	a string
}
type staffs struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func isLoginValid(username string, pass string, table string) (string, bool) {

	var name, username1 string
	row := db.QueryRow("SELECT name,username FROM "+table+" WHERE password=? AND username=?", pass, username)
	e := row.Scan(&name, &username1)
	if e != nil {
		log.Println(e)
		return "INVALID", false
	}
	if username1 == username {
		return name, true
	}
	return "INVALID", false
}
func adminLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	pass := r.FormValue("password")
	var name string
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
func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	pass := r.FormValue("password")
	_, isValid := isLoginValid(username, pass, "staff")
	if isValid {
		fmt.Fprintf(w, "Logged In")
	} else {
		fmt.Fprintf(w, "INVALID")
	}
}
func stuLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	pass := r.FormValue("password")
	name, isValid := isLoginValid(username, pass, "stuLogin")
	if isValid {
		fmt.Fprintf(w, name)
	} else {
		fmt.Fprintf(w, "INVALID")
	}
}
func getQuesHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	tName := r.FormValue("testName")
	qtype := "text"
	fmt.Println(id, "-->", tName)
	var class, ques, op1, op2, op3, op4, ans, level string
	q := "SELECT * FROM " + tName + " WHERE id = ?"
	row := db.QueryRow(q, id)
	e := row.Scan(&id, &class, &ques, &op1, &op2, &op3, &op4, &ans, &level)
	if e != nil {
		fmt.Fprintf(w, "NOT AVAILABLE")
		log.Println(e)
		return
	}
	if strings.Contains(string(op1), "/images/") {
		qtype = "image"
	}
	qobj := &question{ID: id, Class: class, Ques: ques, Option1: op1, Option2: op2, Option3: op3, Option4: op4, Ans: ans, Type: qtype, Level: level}
	qjson, je := json.Marshal(qobj)
	if je != nil {
		log.Print(je)
	}
	fmt.Fprintf(w, string(qjson))
}
func getDataHandler(w http.ResponseWriter, r *http.Request) {
	getTest := r.FormValue("getTest")
	setTest := r.FormValue("setTest")
	testName := r.FormValue("testName")
	status := r.FormValue("status")
	allTest := r.FormValue("allTest")
	if status != "" && testName != "" {
		_, e := db.Exec("update testDetails set status = ? where testName = ?", status, testName)
		if e != nil {
			log.Print(e)
		}
	}
	if getTest == "" && testName != "" && allTest == "" {

		//deleted code
		resJSON, erj := json.Marshal(getAns(testName))
		if erj != nil {
			log.Print(erj)
		}
		fmt.Fprintf(w, string(resJSON))
		return
	} else if setTest == "" {
		if allTest != "" {
			fmt.Fprintf(w, fetchTest(1))
		} else {
			fmt.Fprintf(w, fetchTest(0))
		}
	} else {
		fmt.Fprintf(w, "OK")
	}
}
func createNewTestHandler(w http.ResponseWriter, r *http.Request) {
	testName := r.FormValue("testName")
	no := string(r.FormValue("noOfQues"))
	l := string(r.FormValue("level"))
	if _, err := strconv.Atoi(no + l); err != nil {
		fmt.Fprint(w, "Check number fields")
		return
	}
	levelList := make([]level, 0, 100)
	levels, _ := strconv.Atoi(l)
	for i := 1; i <= levels; i++ {
		temp := level{i, 0}
		levelList = append(levelList, temp)
	}
	jdata, _ := json.Marshal(levelList)

	// TODO: check test already present

	_, e := db.Exec("insert into testDetails values( ? , ? , '', ?, ? )", testName, no, l, string(jdata))
	// code to create the table for the test
	q := "CREATE TABLE " + testName + " ( id int NOT NULL AUTO_INCREMENT, class text, ques text, option1 text, option2 text, option3 text, option4 text,ans text, level text, PRIMARY KEY (id)) "
	db.Exec(q)
	query := "CREATE TABLE IF NOT EXISTS " + testName + "response ( sid varchar(10), qno int, ans varchar(30))"
	_, e1 := db.Exec(query)
	if e != nil || e1 != nil {
		fmt.Fprintf(w, "not OK")
	} else {
		fmt.Fprintf(w, "OK")
	}
}
func saveToTestHandler(w http.ResponseWriter, r *http.Request) {
	testName := r.FormValue("testName")
	id := r.FormValue("id")
	ques := r.FormValue("ques")
	qtype := r.FormValue("type")
	option1 := r.FormValue("option1")
	option2 := r.FormValue("option2")
	option3 := r.FormValue("option3")
	option4 := r.FormValue("option4")
	ans := r.FormValue("ans")
	level := r.FormValue("level")
	fmt.Println(ques, "shit")
	id = strings.TrimSpace(id)
	testName = strings.TrimSpace(testName)
	if qtype == "image" {
		fmt.Println("image type")
		r.ParseMultipartForm(32 << 20)
		file1, handler1, err := r.FormFile("imageoption1")
		if err != nil {
			fmt.Println(err)
			return
		}
		file2, _, _ := r.FormFile("imageoption2")
		file3, _, _ := r.FormFile("imageoption3")
		file4, _, _ := r.FormFile("imageoption4")
		defer file1.Close()
		defer file2.Close()
		defer file3.Close()
		defer file4.Close()
		option1 = "/images/" + testName + id + "option1"
		option2 = "/images/" + testName + id + "option2"
		option3 = "/images/" + testName + id + "option3"
		option4 = "/images/" + testName + id + "option4"
		fmt.Fprintf(w, "%v", handler1.Header)
		f1, err := os.OpenFile("."+option1, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}

		f2, _ := os.OpenFile("."+option2, os.O_WRONLY|os.O_CREATE, 0666)
		f3, _ := os.OpenFile("."+option3, os.O_WRONLY|os.O_CREATE, 0666)
		f4, _ := os.OpenFile("."+option4, os.O_WRONLY|os.O_CREATE, 0666)

		defer f1.Close()
		defer f2.Close()
		defer f3.Close()
		defer f4.Close()
		io.Copy(f1, file1)
		io.Copy(f2, file2)
		io.Copy(f3, file3)
		io.Copy(f4, file4)
	}

	fmt.Println("work akuthu " + ans)
	fmt.Print(testName, "----> ", id)
	rowID := db.QueryRow("SELECT id from "+testName+" WHERE id = ?", id)
	eid := rowID.Scan(&id)
	if eid != nil {
		if eid == sql.ErrNoRows {
			query := "insert into " + testName + "(class, ques, option1, option2, option3, option4, ans, level) values( '' , ?, ? , ? , ? ,? , ?, ? )"
			_, e := db.Exec(query, ques, option1, option2, option3, option4, ans, level)
			if e != nil {
				log.Print(e)
				fmt.Fprintf(w, "error inserting to table"+id)
			} else {

			}
		} else {
			log.Print(eid)
			fmt.Fprintf(w, "error Connecting to table"+id)
		}
	} else {
		query := "UPDATE " + testName + " SET ques=?, option1=?, option2=?, option3=?, option4=?, ans=?, level=? where id = ?"
		_, e := db.Exec(query, ques, option1, option2, option3, option4, ans, level, id)
		if e != nil {
			log.Print(e)
			fmt.Fprintf(w, "error Updating table"+id)
			return
		}
	}

	var nextID int
	row := db.QueryRow("SELECT id from " + testName + " ORDER BY id desc LIMIT 1 ")
	er := row.Scan(&nextID)
	S := (strconv.Itoa(nextID + 1))
	if er != nil {
		fmt.Fprintf(w, " Not done Error occured")
	}
	fmt.Println(S)
	fmt.Fprintf(w, S)

}
func saveResponseHandler(w http.ResponseWriter, r *http.Request) {
	res := r.FormValue("response")
	fmt.Println(res)
	obj := response{}
	json.Unmarshal([]byte(res), &obj)

	tableName := obj.TestName + "response"
	query := "CREATE TABLE IF NOT EXISTS " + tableName + "( sid varchar(10), qno int, ans varchar(30))"
	_, e := db.Exec(query)
	if e != nil {
		fmt.Println("not OK")
		return
	}
	fmt.Println("test1")
	score := 0
	for i, v := range obj.Ans {
		fmt.Println("test2")
		query = "select sid from " + tableName + " where sid = ? and qno = ? LIMIT 1"
		rowID := db.QueryRow(query, obj.Sid, i+1) //check if present
		var temp string
		e = rowID.Scan(&temp)

		if e != nil {
			fmt.Println("goes to err block", e)
			if e == sql.ErrNoRows { //if not present
				fmt.Println("not present block")
				query = "insert into " + tableName + " values (?, ?, ?)"
				_, e1 := db.Exec(query, obj.Sid, i+1, v)
				if e1 != nil {
					fmt.Fprint(w, "not OK", i)
				}
				continue
			} else {
				fmt.Fprint(w, "not OK", i)
			}
		}
		query = "update " + tableName + " set ans = ? where sid = ? and qno = ?" //present so update
		_, e = db.Exec(query, v, obj.Sid, i+1)
		if e != nil {
			fmt.Fprint(w, "not OK", i)
		}
		query = "select ans from " + obj.TestName + " where id = ?"
		var answer string
		row := db.QueryRow(query, i+1)
		e = row.Scan(&answer)
		fmt.Println(v, "--", answer)
		if e == nil {
			fmt.Println("--")
			if answer == v {
				fmt.Println("++")
				score++
			}
		}

	}
	fmt.Println("Score == > ", score)
	fmt.Fprintf(w, strconv.Itoa(score))
}
func showResultHandler(w http.ResponseWriter, r *http.Request) {
	tname := r.FormValue("testName")
	trname := tname + "response"
	stud := make([]string, 0, 100)
	resL := make([]response, 0, 100)
	//ans := getAns(tname);
	var tstud, tans string
	qSud := "select distinct sid from " + trname + " ORDER BY sid ASC;"
	rows, errs := db.Query(qSud)
	if errs != nil {
		log.Print(errs)
		fmt.Println(trname)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&tstud)
		if err != nil {
			log.Print(err)
		}
		stud = append(stud, tstud)
		ans := make([]string, 0, 100)
		//fmt.Println(trname)
		qans := "select ans from " + trname + " where sid = ?"
		rows1, errs1 := db.Query(qans, tstud)
		if errs1 != nil {
			log.Print(errs1)
		}
		defer rows.Close()
		for rows1.Next() {
			err = rows1.Scan(&tans)
			if err != nil {
				log.Print(err)
			}
			ans = append(ans, tans)
		}
		obj := response{TestName: tname, Sid: tstud, Ans: ans}
		resL = append(resL, obj)
	}
	resJSON, erj := json.Marshal(resL)
	if erj != nil {
		log.Print(erj)
	}
	fmt.Fprintf(w, string(resJSON))

}
func changeDataHandler(w http.ResponseWriter, r *http.Request) {
	tname := r.FormValue("testName")
	otype := r.FormValue("type")
	fmt.Print(otype, " called")
	if string(otype) == "rename" {
		newTestName := r.FormValue("newTestName")
		if strings.ContainsAny(newTestName, " ") {
			fmt.Fprintf(w, "No spaces allowed")
			return
		}
		if newTestName != "" && tname != "" {
			renameTest(tname, newTestName)
			fmt.Fprintf(w, "Rename Done")
		}
	}
	if string(otype) == "DELETE" {
		if string(tname) == "" {
			fmt.Fprintf(w, "Test Name Not Received")
			return
		}
		fmt.Println(tname)
		query1 := "DROP TABLE " + tname
		query2 := "DROP TABLE " + tname + "response"
		query3 := "DELETE FROM testDetails where testName= ?"
		_, e1 := db.Exec(query1)
		_, e2 := db.Exec(query2)
		_, e3 := db.Exec(query3, tname)
		if e1 != nil && e2 != nil && e3 != nil {
			fmt.Fprintf(w, "some Error occured")
		} else {
			fmt.Fprintf(w, "Operation Done")
		}
		return

	} else if string(otype) == "status" {
		var status string

		if string(r.FormValue("status")) == "1" {
			status = "ONLINE"
		} else {
			status = "OFFLINE"
		}
		query := "UPDATE testDetails set status = ? where testName = ? "
		fmt.Print(query, status)
		_, e := db.Exec(query, status, tname)
		if e != nil {
			log.Print(e)
		}
		fmt.Fprintf(w, tname+" is "+status)
	}
}
func registerHandler(w http.ResponseWriter, r *http.Request) {
	jsonData := r.FormValue("list")

	var list *[]stuList
	added := make([]string, 0, 100)
	json.Unmarshal([]byte(jsonData), &list)

	for _, v := range *list {

		sql := "INSERT into stuLogin(username, password, name) values(?, ?, ?)"
		if v.Username != "" && v.Password != "" && v.Name != "" {
			_, e := db.Exec(sql, v.Username, v.Password, v.Name)
			if e == nil {
				added = append(added, v.Name)
			}
			fmt.Println(v.Name, "-->", v.Password, "..", v.Username, sql)
		}
	}

	jsonResult, _ := json.Marshal(added)
	fmt.Fprintf(w, string(jsonResult))
}
func changeLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	newUsername := r.FormValue("newUsername")
	oldPassword := r.FormValue("oldPassword")
	newPassword := r.FormValue("newPassword")

	msg := "try again"
	sqlCheckPass := "SELECT password from staff where username = ?"
	row := db.QueryRow(sqlCheckPass, username)

	var tempPass string
	e := row.Scan(&tempPass)
	if e == nil {
		if tempPass == oldPassword {
			var sqlUpdate string
			if newUsername == "" {
				sqlUpdate = "UPDATE staff set password = ? where username = ?"
				fmt.Println("changing password")
				_, e = db.Exec(sqlUpdate, newPassword, username)
			} else {
				fmt.Print("changing username")
				sqlUpdate = "UPDATE staff set username = ? where username = ?"
				_, e = db.Exec(sqlUpdate, newUsername, username)
			}
			if e == nil {
				msg = "Done"
			}
		}
	}
	fmt.Fprintf(w, msg)
}
func levelHandler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println(r.Body)
	if r.FormValue("testName") != "" {
		testName := r.FormValue("testName")
		row := db.QueryRow("select levelJson from testDetails where testName = ?", testName)
		var jsonData string
		e := row.Scan(&jsonData)
		if e != nil {
			log.Println(e)
			return
		}
		//fmt.Println(jsonData)
		fmt.Fprint(w, jsonData)
	} else {

		var t levelSetModel
		var bodyBytes []byte
		bodyBytes, _ = ioutil.ReadAll(r.Body)
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)

		e1 := json.Unmarshal([]byte(bodyString), &t)
		if e1 != nil {
			fmt.Println(e1)
		}
		fmt.Println(t)
		row := db.QueryRow("select level from testDetails where testName = ?", t.TestName)
		var temp string
		e := row.Scan(&temp)
		if e != nil {
			log.Println(e)
			return
		}
		lev, _ := strconv.Atoi(temp)
		if lev != len(t.Levels) {
			log.Println("Error length of the levelJson and levels not equal")
			fmt.Println(lev, len(t.Levels))
			return
		}
		jdata, _ := json.Marshal(t.Levels)

		db.Exec("update testDetails set levelJson = ? where testName =?", string(jdata), t.TestName)
	}
}
func getAns(testName string) []ansLevel {

	fmt.Println(testName)
	qAns := "select ans, level from " + testName
	rows, errs := db.Query(qAns)
	if errs != nil {
		log.Print(errs)
	}
	defer rows.Close()
	var ta, tl string
	objList := make([]ansLevel, 0, 100)
	for rows.Next() {
		err := rows.Scan(&ta, &tl)
		if err != nil {
			log.Print(err)
		}
		obj := &ansLevel{Ans: ta, Level: tl}
		objList = append(objList, *obj)
	}
	return objList
}
func fetchTest(c int) string {
	//fmt.Println("a")
	var name, no, level, status string
	objArr := make([]testList, 0, 10)
	var q string
	if c == 1 {
		q = "select testName, noOfQues, level, status from testDetails"
	} else {
		q = "select testName, noOfQues, level, status from testDetails where status = 'ONLINE'"
	}
	rows, errs := db.Query(q)
	if errs != nil {
		log.Print(errs)
	}
	if rows == nil {
		return "not available"
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &no, &level, &status)
		tobj := &testList{Name: name, No: no, Level: level, Status: status}
		objArr = append(objArr, *tobj)

		if err != nil {
			log.Print(err)
		}
		fmt.Println(name)
	}
	tjson, je := json.Marshal(objArr)
	if je != nil {
		log.Print(je)
	}

	return string(tjson)
}
func renameTest(oldName string, newName string) {
	//fmt.Println("a")
	q1 := "RENAME TABLE " + oldName + " TO " + newName
	q2 := "RENAME TABLE " + oldName + "response TO " + newName + "response"
	q3 := "update testDetails set testName = '" + newName + "' where testName = '" + oldName + "'"
	_, e1 := db.Exec(q1)
	_, e2 := db.Exec(q2)
	_, e3 := db.Exec(q3)
	if e1 != nil || e2 != nil || e3 != nil {
		log.Print(e1, e2, e3)
	}
}
func fetchTestData(testName string) string {

	var name, no string
	result := ""
	rows, errs := db.Query("select testName,noOfQues from testDetails where testName = ?", testName)
	if errs != nil {
		log.Print(errs)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &no)
		if err != nil {
			log.Print(err)
		}
		result = result + name + "\t" + no
	}
	return result
}

func main() {
	db, _ = sql.Open("mysql", user+":"+password+"@/"+database)
	if err := db.Ping(); err != nil {
		log.Print(err)
	}
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.HandleFunc("/adminLogin/", adminLoginHandler)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/stuLogin/", stuLoginHandler)
	http.HandleFunc("/getQuestions/", getQuesHandler)
	http.HandleFunc("/getData/", getDataHandler)
	http.HandleFunc("/createNewTest/", createNewTestHandler)
	http.HandleFunc("/saveToTest/", saveToTestHandler)
	http.HandleFunc("/saveResponse/", saveResponseHandler)
	http.HandleFunc("/showResult/", showResultHandler)
	http.HandleFunc("/changeData/", changeDataHandler)
	http.HandleFunc("/register/", registerHandler)
	http.HandleFunc("/changeLogin/", changeLoginHandler)
	http.HandleFunc("/level/", levelHandler)
	fmt.Print("Started Server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
