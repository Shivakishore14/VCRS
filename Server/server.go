package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
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

type question struct {
	Id      string `json:"id"`
	Class   string `json:"class"`
	Ques    string `json:"question"`
	Option1 string `json:"option1"`
	Option2 string `json:"option2"`
	Option3 string `json:"option3"`
	Option4 string `json:"option4"`
	Ans     string `json:"ans"`
	Type    string `json:"type"`
}
type testList struct {
	Name string `json:"name"`
	No   string `json:"no"`
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

func isLoginValid(username string, pass string) bool {
	db, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err = db.Ping(); err != nil {
		log.Print(err)
		return false
	}
	var name string
	defer db.Close()
	row := db.QueryRow("SELECT username FROM staff WHERE password=? AND username=?", pass, username)
	e := row.Scan(&name)
	if e != nil {
		log.Println(e)
		return false
	}
	if name == username {
		return true
	}
	return false
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	pass := r.FormValue("password")
	isValid := isLoginValid(username, pass)
	if isValid {
		fmt.Fprintf(w, "Logged In")
	} else {
		fmt.Fprintf(w, "INVALID")
	}
}
func getQuesHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	tName := r.FormValue("testName")
	qtype := "text"
	fmt.Println(id, "-->", tName)
	db, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err = db.Ping(); err != nil {
		log.Print(err)
	}
	var class, ques, op1, op2, op3, op4, ans string
	defer db.Close()
	q := "SELECT * FROM " + tName + " WHERE id = ?"
	row := db.QueryRow(q, id)
	e := row.Scan(&id, &class, &ques, &op1, &op2, &op3, &op4, &ans)
	if e != nil {
		fmt.Fprintf(w, "NOT AVAILABLE")
		log.Println(e)
		return
	}
	if strings.Contains(string(op1), "/images/") {
		qtype = "image"
	}
	qobj := &question{Id: id, Class: class, Ques: ques, Option1: op1, Option2: op2, Option3: op3, Option4: op4, Ans: ans, Type: qtype}
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
		db, err := sql.Open("mysql", user+":"+password+"@/"+database)
		if err = db.Ping(); err != nil {
			log.Print(err)
		}
		defer db.Close()
		_, e := db.Exec("update testDetails set status = ? where testName = ?", status, testName)
		if e != nil {
			log.Print(e)
		}
	}
	if getTest == "" && testName != "" && allTest == "" {

		//deleted code
		resJson, erj := json.Marshal(getAns(testName))
		if erj != nil {
			log.Print(erj)
		}
		fmt.Fprintf(w, string(resJson))
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
	no := r.FormValue("noOfQues")
	db, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err = db.Ping(); err != nil {
		log.Print(err)
	}
	defer db.Close()
	_, e := db.Exec("insert into testDetails values( ? , ? , '', '' )", testName, no)
	// code to create the table for the test
	q := "CREATE TABLE " + testName + " ( id int NOT NULL AUTO_INCREMENT, class text, ques text, option1 text, option2 text, option3 text, option4 text,ans text, PRIMARY KEY (id)) "
	_, e1 := db.Exec(q)
	query := "CREATE TABLE IF NOT EXISTS " + testName + "response ( sid varchar(10), qno int, ans varchar(30))"
	_, e1 = db.Exec(query)
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
	db, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err = db.Ping(); err != nil {
		log.Print(err)
	}
	defer db.Close()
	fmt.Print(testName, "----> ", id)
	rowId := db.QueryRow("SELECT id from "+testName+" WHERE id = ?", id)
	eid := rowId.Scan(&id)
	if eid != nil {
		if eid == sql.ErrNoRows {
			query := "insert into " + testName + "(class, ques, option1, option2, option3, option4, ans) values( '' , ?, ? , ? , ? ,? , ? )"
			_, e := db.Exec(query, ques, option1, option2, option3, option4, ans)
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
		query := "UPDATE " + testName + " SET ques=?, option1=?, option2=?, option3=?, option4=?, ans=? where id = ?"
		_, e := db.Exec(query, ques, option1, option2, option3, option4, ans, id)
		if e != nil {
			log.Print(e)
			fmt.Fprintf(w, "error Updating table"+id)
			return
		}
	}

	var nextId int
	row := db.QueryRow("SELECT id from " + testName + " ORDER BY id desc LIMIT 1 ")
	er := row.Scan(&nextId)
	S := (strconv.Itoa(nextId + 1))
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
	db, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err = db.Ping(); err != nil {
		log.Print(err)
	}
	defer db.Close()
	_, e := db.Exec(query)
	if e != nil {
		fmt.Println("not OK")
		return
	}
	for i, v := range obj.Ans {
		query = "select * from " + tableName + "where sid = ? and qno = ?"
		_, e = db.Exec(query, obj.Sid, i+1) //check if present
		if e != nil {
			if e == sql.ErrNoRows { //if not present
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
	}
	fmt.Fprintf(w, "OK")
}
func showResultHandler(w http.ResponseWriter, r *http.Request) {
	tname := r.FormValue("testName")
	trname := tname + "response"
	db, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err = db.Ping(); err != nil {
		log.Print(err, tname)
	}
	defer db.Close()
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
	resJson, erj := json.Marshal(resL)
	if erj != nil {
		log.Print(erj)
	}
	fmt.Fprintf(w, string(resJson))

}
func changeDataHandler(w http.ResponseWriter, r *http.Request) {
	tname := r.FormValue("testName")
	otype := r.FormValue("type")
	db, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err = db.Ping(); err != nil {
		log.Print(err)
	}
	defer db.Close()
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
			fmt.Fprintf(w, "Operation Done")
		} else {
			fmt.Fprintf(w, "some Error occured")
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
func getAns(testName string) []string {
	db, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err = db.Ping(); err != nil {
		log.Print(err)
	}
	defer db.Close()
	fmt.Println(testName)
	qAns := "select ans from " + testName
	rows, errs := db.Query(qAns)
	if errs != nil {
		log.Print(errs)
	}
	defer rows.Close()
	var ta string
	ans := make([]string, 0, 100)
	for rows.Next() {
		err := rows.Scan(&ta)
		if err != nil {
			log.Print(err)
		}
		ans = append(ans, ta)
	}
	return ans
}
func fetchTest(c int) string {
	//fmt.Println("a")
	db, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err = db.Ping(); err != nil {
		log.Print(err)
	}
	defer db.Close()
	var name, no string
	objArr := make([]testList, 0, 10)
	q := ""
	if c == 1 {
		q = "select testName, noOfQues from testDetails"
	} else {
		q = "select testName, noOfQues from testDetails where status = 'ONLINE'"
	}
	rows, errs := db.Query(q)
	if errs != nil {
		log.Print(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &no)
		tobj := &testList{Name: name, No: no}
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
	db, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err = db.Ping(); err != nil {
		log.Print(err)
	}
	defer db.Close()
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
	db, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err = db.Ping(); err != nil {
		log.Print(err)
	}
	defer db.Close()
	var name, no string
	result := ""
	rows, errs := db.Query("select testName,noOfQues from testDetails where testName = ?", testName)
	if errs != nil {
		log.Print(err)
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
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/getQuestions/", getQuesHandler)
	http.HandleFunc("/getData/", getDataHandler)
	http.HandleFunc("/createNewTest/", createNewTestHandler)
	http.HandleFunc("/saveToTest/", saveToTestHandler)
	http.HandleFunc("/saveResponse/", saveResponseHandler)
	http.HandleFunc("/showResult/", showResultHandler)
	http.HandleFunc("/changeData/", changeDataHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
