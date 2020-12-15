package mongoclient

import (
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func jsonToBsonM(theJSON []byte) (bson.M, error) {
	var result bson.M

	errConv := json.Unmarshal(theJSON, &result)
	return result, errConv
}

func jsonToBsonD(theJSON []byte) (bson.D, error) {
	var result bson.D

	//errConv := json.Unmarshal(pJSON, &result)
	buf, errConv := bson.Marshal(theJSON)
	log.Println("buf:", buf)
	return result, errConv
}
