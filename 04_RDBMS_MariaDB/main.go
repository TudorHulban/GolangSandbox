package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "hera:hera@tcp(localhost:3306)/devops")
	checkErr(err)
	defer db.Close()

	//create table if not exists
	_, err = db.Exec("CREATE TABLE if not exists userinfo ( `uid` INTEGER PRIMARY KEY AUTOINCREMENT,`username` VARCHAR(64) NULL,`departname` VARCHAR(64) NULL,`created` DATE NULL )")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("john", "HR", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("LastInsertId: ", id)

	//query row
	var uid int
	var username string
	var department string
	var created time.Time

	_ = db.QueryRow("SELECT * FROM userinfo where username=?", "john").Scan(&uid, &username, &department, &created)

	fmt.Println("row ID: ", uid)
	fmt.Println("row name: ", username)
	fmt.Println("row department: ", department)
	fmt.Println("row created: ", created)

	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("john smith", id)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("RowsAffected: ", affected)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	defer rows.Close() //good habit to close after error check

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)

		fmt.Println("user ID: ", uid)
		fmt.Println("user name: ", username)
		fmt.Println("department: ", department)
		fmt.Println("created: ", created)
	}

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affected, err = res.RowsAffected()
	checkErr(err)

	fmt.Println("RowsAffected: ", affected)
}

func checkErr(err error) {
	if err != nil {
		log.Print("got: ", err)
		os.Exit(1)
	}
}
