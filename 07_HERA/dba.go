package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

// DBSQLiteInfo Struct to use for defining methods for satisfying interface.
type DBSQLiteInfo struct {
	pathSQLiteFile string
}

// DBPostgresInfo Struct to use for defining methods for satisfying interface.
type DBPostgresInfo struct {
	port     uint
	ip       string
	user     string
	password string
}

// DBMariaInfo Struct to use for defining methods for satisfying interface.
type DBMariaInfo struct {
	port     uint
	ip       string
	user     string
	password string
}

// RDBMS Interface for interacting with RDBMS.
type RDBMS interface {
	Open(dbName string) (*sql.DB, error)
	TableExists(db *sql.DB, dbName, tableName string) bool
	CreateTable(db *sql.DB, dbName, tableName, ddl string, columnPKAutoincrement bool) bool
	SingleInsert(db *sql.DB, tableName string, pValues []string) error
	BulkInsert(db *sql.DB, tableName string, columnNames []string, pValues [][]string) error
}

// Open Method opens database.
func (lite DBSQLiteInfo) Open(dbName string) (*sql.DB, error) {
	return sql.Open("sqlite3", lite.pathSQLiteFile)
}

// TableExists Method checks if database table exists.
func (lite DBSQLiteInfo) TableExists(db *sql.DB, dbName, tableName string) bool {
	var occurences int
	_ = db.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?", tableName).Scan(&occurences)

	return (occurences == 1)
}

func (lite DBSQLiteInfo) CreateTable(db *sql.DB, dbName, tableName, ddl string, columnPKAutoincrement bool) bool {
	var theDDL string

	if columnPKAutoincrement {
		theDDL = "\"id\" INTEGER PRIMARY KEY AUTOINCREMENT," + ddl
	} else {
		theDDL = ddl
	}

	theDDL = "CREATE TABLE " + tableName + "(" + ddl + ")"

	_, err := db.Exec(theDDL)
	errors.WithMessage(err, "DB SQLite CreateTable")

	return lite.TableExists(db, dbName, tableName)
}

func (pg DBPostgresInfo) Open(dbName string) (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pg.ip, pg.port, pg.user, pg.password, dbName)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		fmt.Println("DB Postgres Open: ", err)
		return nil, err
	}

	return db, db.Ping()
}

func (maria DBMariaInfo) Open(dbName string) (*sql.DB, error) {
	dbinfo := maria.user + ":" + maria.password + "@tcp(" + maria.ip + ":" + strconv.Itoa(maria.port) + ")/" + dbName
	return sql.Open("mysql", dbinfo)
}

func (rDBPostgresInfo DBPostgresInfo) TableExists(db *sql.DB, dbName, tableName string) bool {
	var occurences bool

	theDML := "SELECT exists (select 1 from information_schema.tables WHERE table_schema='public' AND table_name=" + "'" + tableName + "'" + ")"
	_ = db.QueryRow(theDML).Scan(&occurences)

	return occurences
}

func (maria DBMariaInfo) TableExists(db *sql.DB, dbName, tableName string) bool {
	var occurences bool

	theDML := "select count(1) from information_schema.tables WHERE table_schema=" + "'" + dbName + "'" + " AND table_name=" + "'" + tableName + "'" + " limit 1"
	_ = db.QueryRow(theDML).Scan(&occurences)

	return occurences
}

func (pg DBPostgresInfo) CreateTable(db *sql.DB, dbName, tableName, ddl string, columnPKAutoincrement bool) bool {
	var theDDL string

	if columnPKAutoincrement {
		theDDL = "\"id\" serial," + ddl
	} else {
		theDDL = ddl
	}

	theDDL = "CREATE TABLE " + tableName + "(" + theDDL + ")"

	_, err := db.Exec(theDDL)
	return errors.WithMessage(err, "DB Postgres CreateTable")

	return pg.TableExists(db, dbName, tableName)
}

func (rDBMariaInfo DBMariaInfo) CreateTable(db *sql.DB, dbName, tableName, ddl string, columnPKAutoincrement bool) bool {
	var theDDL string

	if columnPKAutoincrement {
		theDDL = "\"id\" serial," + ddl
	} else {
		theDDL = ddl
	}

	theDDL = "CREATE TABLE " + tableName + " (" + strings.Replace(theDDL, "\"", "", -1) + ")"

	fmt.Println(theDDL, columnPKAutoincrement)

	_, err := db.Exec(theDDL)
	checkErr(err, "rDBMariaInfo CreateTable: "+theDDL)

	return rDBMariaInfo.TableExists(db, dbName, tableName)
}

func (rDBSQLiteInfo DBSQLiteInfo) SingleInsert(db *sql.DB, tableName string, values []string) error {
	theDDL := "insert into " + tableName + " values(" + "\"" + strings.Join(values, "\""+","+"\"") + "\"" + ")"
	_, err := db.Exec(theDDL)

	return err
}

func (rDBPostgresInfo DBPostgresInfo) SingleInsert(db *sql.DB, tableName string, values []string) error {
	theDDL := "insert into " + tableName + " values(" + "\"" + strings.Join(values, "\""+","+"\"") + "\"" + ")"
	_, err := db.Exec(theDDL)

	return err
}

func (rDBMariaInfo DBMariaInfo) SingleInsert(db *sql.DB, tableName string, values []string) error {
	theDDL := "insert into " + tableName + " values(" + "\"" + strings.Join(values, "\""+","+"\"") + "\"" + ")"
	_, err := db.Exec(theDDL)

	return err
}

func (rDBSQLiteInfo DBSQLiteInfo) BulkInsert(db *sql.DB, tableName string, pColumnNames []string, pValues [][]string) error {
	theQuestionMarks := returnNoValues(pValues[0], "?")

	// -------- DB Transaction Start -----------
	dbTransaction, err := db.Begin()
	checkErr(err, "db.Begin")

	statement := "insert into " + tableName + "(" + strings.Join(pColumnNames, ",") + ")" + " values " + theQuestionMarks

	dml, err := dbTransaction.Prepare(statement)
	checkErr(err, "dbTransaction.Prepare")
	defer dml.Close()

	for _, columnValues := range pValues {
		_, err := dml.Exec(SliceToInterface(columnValues)...)

		if err != nil {
			fmt.Println("dml.Exec: ", err)
			panic(err)
		}
	}
	dbTransaction.Commit()
	// -------- DB Transaction End -----------

	return err
}

func (rDBMariaInfo DBMariaInfo) BulkInsert(db *sql.DB, tableName string, pColumnNames []string, pValues [][]string) error {
	theQuestionMarks := returnNoValues(pValues[0], "?")

	// -------- DB Transaction Start -----------
	dbTransaction, err := db.Begin()
	checkErr(err, "db.Begin")

	statement := "insert into " + tableName + "(" + strings.Join(pColumnNames, ",") + ")" + " values " + theQuestionMarks

	dml, err := dbTransaction.Prepare(statement)
	checkErr(err, "dbTransaction.Prepare")
	defer dml.Close()

	for _, columnValues := range pValues {
		_, err := dml.Exec(SliceToInterface(columnValues)...)

		if err != nil {
			fmt.Println("dml.Exec: ", err)
			panic(err)
		}
	}
	dbTransaction.Commit()
	// -------- DB Transaction End -----------

	return err
}

func prepareDBField(field string) string {
	return strings.Replace(field, "'", "\"", -1)
}

func returnNoValues(inputSlice []string, charToReturn string) string {
	var toReturn string
	for range inputSlice {
		toReturn = toReturn + charToReturn + ","
	}

	return "(" + toReturn[0:len(toReturn)-1] + ")"
}

func wrapSliceValuesx(inputSlice []string, charToWrap string) string {
	return "(" + charToWrap + strings.Join(inputSlice, charToWrap+","+charToWrap) + charToWrap + ")"
}

func SliceToInterface(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	result := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		result[i] = s.Index(i).Interface()
	}

	return result
}
