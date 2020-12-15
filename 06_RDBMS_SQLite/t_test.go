package main

import (
	"testing"
)

func TestSQLiteOpers(t *testing.T) {
	/*
		//check table exists
		log.Println("tbExists: ", tableExists(db, "userinfo"))

		//create table version
		if !tableExists(db, "version") {
			if errCreate := createTable(db, "version", "(`name` text, `date` text, `description` text)"); errCreate != nil {
				log.Fatal(errCreate)
			}
		}

		//create table if not exists
		_, err = db.Exec("CREATE TABLE if not exists userinfo ( `uid` INTEGER PRIMARY KEY AUTOINCREMENT,`username` VARCHAR(64) NULL,`departname` VARCHAR(64) NULL,`created` DATE NULL )")
		if err != nil {
			log.Fatal(err)
		}

		// insert
		stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec("john", "HR", "2012-12-09")
		if err != nil {
			log.Fatal(err)
		}

		id, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("LastInsertId: ", id)

		//query row
		var uid int
		var username string
		var department string
		var created time.Time

		_ = db.QueryRow("SELECT * FROM userinfo where username=?", "john").Scan(&uid, &username, &department, &created)

		log.Println("row ID: ", uid)
		log.Println("row name: ", username)
		log.Println("row department: ", department)
		log.Println("row created: ", created)

		//transform to json
		json, err := getAsJSON(db, "SELECT * FROM userinfo")
		if err != nil {
			log.Fatal(err)
		}

		log.Println("getJSON data: ", json)

		//check department
		log.Println("is Exact: ", valueExact(json, "departname"))

		// update
		stmt, err = db.Prepare("update userinfo set username=? where uid=?")
		if err != nil {
			log.Fatal(err)
		}

		res, err = stmt.Exec("john smith", id)
		if err != nil {
			log.Fatal(err)
		}

		affected, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("RowsAffected: ", affected)

		// query
		rows, err := db.Query("SELECT * FROM userinfo")
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close() //good habit to close

		for rows.Next() {
			err = rows.Scan(&uid, &username, &department, &created)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("user name: %v, ID: %v, created: %v, department: %v", username, uid, created, department)
		}

		// delete
		stmt, err = db.Prepare("delete from userinfo where uid=?")
		if err != nil {
			log.Fatal(err)
		}

		res, err = stmt.Exec(id)
		if err != nil {
			log.Fatal(err)
		}

		affected, err = res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("RowsAffected: ", affected)
		db.Close()
	*/
}
