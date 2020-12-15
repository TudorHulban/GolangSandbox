package mongoclient

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/TudorHulban/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	john = moValue{
		Name:   "john",
		Gender: "male",
		Age:    44,
	}

	mary = moValue{
		Name:   "mary",
		Gender: "female",
		Age:    44,
	}
)

func testCfg() *Cfg {
	return &Cfg{
		TimeoutSecs: 3,
		URL:         "mongodb://localhost:27017",
		Database:    "testing",
		Collection:  "persons",
		l:           log.New(log.DEBUG, os.Stderr, true),
	}
}

// testInsertOne Helper for inserting one value.
func testInsertOne(ctx context.Context, t *testing.T, m *MoInstance, v moValue) primitive.ObjectID {
	j, errMarshall := json.Marshal(v)
	require.Nil(t, errMarshall)

	id, errInsert := m.InsertOne(ctx, j)
	require.Nil(t, errInsert)
	//assert.Greater(t, id, 0)

	m.l.Printf("inserted ID: %s", id.String())
	return id
}

func TestInsertOne(t *testing.T) {
	m, errNew := NewMongo(testCfg())
	require.Nil(t, errNew, "connection to Mongo DB issues")
	require.NotNil(t, m)

	ctx := context.Background()
	require.Nil(t, m.Connect(ctx), "could not connect")
	defer m.Disconnect(ctx)

	m.l.Debug("Prepare to insert john")
	johnID := testInsertOne(ctx, t, m, john)
	m.l.Print("johnID: ", johnID)
}

func TestFindByID(t *testing.T) {
	m, errNew := NewMongo(testCfg())
	require.Nil(t, errNew, "connection to Mongo DB issues")
	require.NotNil(t, m)

	ctx := context.Background()
	require.Nil(t, m.Connect(ctx), "could not connect")
	defer m.Disconnect(ctx)

	m.l.Debug("Prepare to insert john")
	johnID := testInsertOne(ctx, t, m, john)
	m.l.Print("johnID: ", johnID)

	record, errFindID := m.FindByID(ctx, johnID)
	require.Nil(t, errFindID, "error find by ID")
	require.NotNil(t, record)
	m.l.Printf("--------- Found (1): %s", record)
}

// TestFindOneRecord Should find a record.
func TestFindOneRecord(t *testing.T) {
	m, errNew := NewMongo(testCfg())
	require.Nil(t, errNew, "connection to Mongo DB issues")
	require.NotNil(t, m)

	ctx := context.Background()
	require.Nil(t, m.Connect(ctx), "could not connect")
	defer m.Disconnect(ctx)

	jsonJohn, errMarshall := json.Marshal(john)
	require.Nil(t, errMarshall)

	record, errFind := m.FindOne(ctx, jsonJohn)
	require.Nil(t, errFind)
	m.l.Printf("--------- Found (1): %s", record)
}

// TestFindOneNoRecord Should not find a record.
func TestFindOneNoRecord(t *testing.T) {
	m, errNew := NewMongo(testCfg())
	require.Nil(t, errNew, "connection to Mongo DB issues")
	require.NotNil(t, m)

	ctx := context.Background()
	require.Nil(t, m.Connect(ctx), "could not connect")
	defer m.Disconnect(ctx)

	// TODO: implementation.
}

// TestFindMany Should find several records.
func TestFindManyJSON(t *testing.T) {
	m, errNew := NewMongo(testCfg())
	require.Nil(t, errNew, "connection to Mongo DB issues")
	require.NotNil(t, m)

	ctx := context.Background()
	require.Nil(t, m.Connect(ctx), "could not connect")
	defer m.Disconnect(ctx)

	testInsertOne(ctx, t, m, mary)

	age := struct{ age uint }{
		age: 44,
	}
	jsonAge, errMarshall := json.Marshal(age)
	require.Nil(t, errMarshall)

	record, errMany := m.FindManyFilterJSON(ctx, jsonAge)
	require.Nil(t, errMany)
	assert.Greater(t, len(record), 2)

	m.l.Printf("--------- Found (%d):", len(record))
	for k, v := range record {
		m.l.Print(k, v)
	}
}

// TestDeleteOne Should delete one record provided the passed filter.
func TestDeleteOne(t *testing.T) {
	m, errNew := NewMongo(testCfg())
	require.Nil(t, errNew, "connection to Mongo DB issues")
	require.NotNil(t, m)

	ctx := context.Background()
	require.Nil(t, m.Connect(ctx), "could not connect")
	defer m.Disconnect(ctx)

	jsonMary, errMarshall := json.Marshal(mary)
	require.Nil(t, errMarshall)

	manyMaryBefore, errMany := m.FindManyFilterJSON(ctx, jsonMary)

	_, errDelete := m.DeleteOne(ctx, jsonMary)
	require.Nil(t, errDelete)

	manyMaryAfter, errMany := m.FindManyFilterJSON(ctx, jsonMary)
	require.Nil(t, errMany)
	assert.Equal(t, len(manyMaryAfter), len(manyMaryBefore)-1)
}

// TestDeleteAll Should delete all records for passed filter.
func TestDeleteAll(t *testing.T) {
	m, errNew := NewMongo(testCfg())
	require.Nil(t, errNew, "connection to Mongo DB issues")
	require.NotNil(t, m)

	ctx := context.Background()
	require.Nil(t, m.Connect(ctx), "could not connect")
	defer m.Disconnect(ctx)

	raw, errMarshall := json.Marshal(john)
	require.Nil(t, errMarshall)

	_, errDelete := m.DeleteAll(ctx, raw)
	require.Nil(t, errDelete)

	many, errMany := m.FindManyFilterJSON(ctx, raw)
	require.Nil(t, errMany)
	assert.Equal(t, len(many), 0)
}

// TestUpdateByID Should update the record with passed ID.
func TestUpdateByID(t *testing.T) {
	m, errNew := NewMongo(testCfg())
	require.Nil(t, errNew, "connection to Mongo DB issues")
	require.NotNil(t, m)

	ctx := context.Background()
	require.Nil(t, m.Connect(ctx), "could not connect")
	defer m.Disconnect(ctx)

	id := testInsertOne(ctx, t, m, mary)
	bsonUpdate := bson.M{"$set": bson.M{"Name": "mary", "Gender": "female", "Age": 45}} // just changing the age

	_, errUpdate := m.UpdateByID(ctx, id, bsonUpdate)
	require.Nil(t, errUpdate)

	maryUpdated := moValue{
		Name:   "mary",
		Gender: "female",
		Age:    45,
	}

	raw, errMarshall := json.Marshal(maryUpdated)
	require.Nil(t, errMarshall)

	record, errFind := m.FindOne(ctx, raw)
	require.Nil(t, errFind)
	require.NotNil(t, record)
	m.l.Printf("record: %s", record)
}

// TestUpdateOne Should update a record given passed data.
func TestUpdateOne(t *testing.T) {
	m, errNew := NewMongo(testCfg())
	require.Nil(t, errNew, "connection to Mongo DB issues")
	require.NotNil(t, m)

	ctx := context.Background()
	require.Nil(t, m.Connect(ctx), "could not connect")
	defer m.Disconnect(ctx)

	testInsertOne(ctx, t, m, mary)

	// define filter for the value to be updated
	bsonFilter := bson.M{
		"name": bson.M{
			"$eq": "mary",
		},
	}

	// define data to update with
	bsonUpdate := bson.M{
		"$set": bson.M{
			"age": 55,
		},
	}

	_, errUpdate := m.UpdateOne(ctx, bsonFilter, bsonUpdate)
	require.Nil(t, errUpdate)
}
