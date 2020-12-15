package main

import (
	"encoding/json"
	"log"
)

type updateFieldValues struct {
	clauseColumn string
	replaceWith  string
}

//JSON4Update - struct implementing receiver for interface
type JSON4Update struct{}

//JSON2Update - interface that can be used with raw JSON to update database table
type JSON2Update interface {
	Update(serverInfo *DBServerInfo, rawJSONData []byte) (int64, error)
}

//Update - method, signature satisfies JSON2Update interface, takes database server info and raw JSON and returns rows updated and error
func (receiver *JSON4Update) Update(serverInfo *DBServerInfo, rawJSONData []byte) (int64, error) {

	var jsonAbstract interface{}
	err := json.Unmarshal(rawJSONData, &jsonAbstract)
	if err != nil {
		log.Fatal("json.Unmarshal:", err)
	}

	log.Println("jsonAbstract:", jsonAbstract)

	jsonLayer1 := jsonAbstract.(map[string]interface{})

	log.Println("jsonLayer1:", jsonLayer1)

	var databaseName string
	var tableName string

	var columnTarget string
	var columnClause string

	var injectPreparedStatement []updateFieldValues

	for keyLayer1, valueLayer1 := range jsonLayer1 {
		log.Println("keyLayer1:", keyLayer1)
		log.Println("valueLayer1:", valueLayer1)

		switch layer1Value := valueLayer1.(type) {

		case string:
			{
				switch keyLayer1 {
				case "Database":
					{
						log.Println("Database:", layer1Value)
						databaseName = layer1Value
					}
				case "Table":
					{
						log.Println("Table:", layer1Value)
						tableName = layer1Value
					}
				}

			}

		case []interface{}:
			{
				log.Println("Data:", layer1Value)

				for keyLayer2, valueLayer2 := range layer1Value {
					log.Println("keyLayer2:", keyLayer2)
					log.Println("valueLayer2:", valueLayer2) //valueLayer2: map[Update:map[name:Tudor] Where:map[id:1]]

					var injectValues updateFieldValues

					for keyLayer3, valueLayer3 := range valueLayer2.(map[string]interface{}) {
						log.Println("--keyLayer3:", keyLayer3)
						log.Println("--valueLayer3:", valueLayer3)

						switch keyLayer3 {

						case "Update":
							{
								log.Println("keyLayer3 - Update:", keyLayer3)

								for keyLayer4, valueLayer4 := range valueLayer3.(map[string]interface{}) {
									log.Println("keyLayer4 - Update:", keyLayer4)
									columnTarget = keyLayer4
									log.Println("valueLayer4 - Update:", valueLayer4)
									injectValues.replaceWith = valueLayer4.(string)
									log.Println(keyLayer4, "injectValues.replaceWith:", injectValues.replaceWith)
								}
							}
						case "Where":
							{
								log.Println("keyLayer3 - Where:", keyLayer3)

								for keyLayer4, valueLayer4 := range valueLayer3.(map[string]interface{}) {
									log.Println("keyLayer4 - Where:", keyLayer4)
									columnClause = keyLayer4
									log.Println("valueLayer4 - Where:", valueLayer4)
									injectValues.clauseColumn = valueLayer4.(string)
									log.Println(keyLayer4, "injectValues.clauseColumn:", injectValues.clauseColumn)
								}

							}
						}

					}

					log.Println(keyLayer2, "-------- injectValues:", injectValues)
					injectPreparedStatement = append(injectPreparedStatement, injectValues)

				}
			}
		}
	}

	log.Println("no values:", len(injectPreparedStatement), "injectPreparedStatement:", injectPreparedStatement)

	db, err := new(DBHandler).Connect(serverInfo)
	if err != nil {
		log.Fatal("new(DBHandler).Connect:", err)
	}

	dbTransaction, err := db.Begin()
	if err != nil {
		log.Fatal("db.Begin:", err)
	}

	statement := "update " + databaseName + "." + tableName + " set " + columnTarget + "=?" + " where " + columnClause + "=?"

	dml, err := dbTransaction.Prepare(statement)
	if err != nil {
		log.Fatal("db.Prepare:", err)
	}
	defer dml.Close()

	log.Println("dml:", dml)

	var rowCount int64

	for key := range injectPreparedStatement {

		log.Println(injectPreparedStatement[key].clauseColumn, injectPreparedStatement[key].replaceWith)
		log.Println("update " + databaseName + "." + tableName + " set " + columnTarget + "=" + injectPreparedStatement[key].replaceWith + " where " + columnClause + "=" + injectPreparedStatement[key].clauseColumn)

		sqlResult, err := dml.Exec(injectPreparedStatement[key].replaceWith, injectPreparedStatement[key].clauseColumn)
		if err != nil {
			log.Fatal("dml.Exec:", err)
		}

		log.Println("sqlResult:", sqlResult)

		updated, err := sqlResult.RowsAffected()
		if err != nil {
			dbTransaction.Rollback()
			log.Fatal("sqlResult.RowsAffected:", err)
			break
		}

		rowCount = rowCount + updated
	}

	dbTransaction.Commit()

	return rowCount, nil
}
