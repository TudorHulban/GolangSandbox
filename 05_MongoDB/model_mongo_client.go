package mongoclient

import (
	"context"
	"time"

	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// File creates Mongo client model.
// Local context timeouts use global cfg but could be passed as argument if neded.
// Returning primitive.ObjectID which is a byte array.
// TODO: verify objID, _ := primitive.ObjectIDFromHex(id) transformation

type Cfg struct {
	TimeoutSecs uint
	URL         string
	Database    string
	Collection  string
	l           *log.LogInfo
}

type MoInstance struct {
	*Cfg

	client *mongo.Client
}

// moValue Data structure for loading into the database.
type moValue struct {
	Age    uint
	Name   string
	Gender string
}

// NewMongo Constructor for Mongo client.
// Caller would need to handle connect / disconnect.
func NewMongo(c *Cfg) (*MoInstance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.TimeoutSecs)*time.Second)
	defer cancel()

	instance, errConnect := mongo.Connect(ctx, options.Client().ApplyURI(c.URL))
	if errConnect != nil {
		return nil, errConnect
	}

	if errPing := instance.Ping(ctx, readpref.Primary()); errPing != nil {
		c.l.Debug("database did not respond")
		return nil, errPing
	}

	c.l.Debug("database did respond to ping")

	result, _ := mongo.NewClient(options.Client().ApplyURI(c.URL))
	if result == nil {
		c.l.Debug("result is nil")
		return nil, errors.New("?")
	}

	return &MoInstance{
		Cfg:    c,
		client: result,
	}, nil
}

// Connect Method connects client instance to configured database.
func (m *MoInstance) Connect(ctx context.Context) error {
	return m.client.Connect(ctx)
}

// Disconnect Method disconnects client from database.
func (m *MoInstance) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

// InsertOne Method inserts the data and returns the ID of the inserted data and error.
func (m *MoInstance) InsertOne(ctx context.Context, data []byte) (primitive.ObjectID, error) {
	ctxLocal, cancel := context.WithTimeout(ctx, time.Duration(m.TimeoutSecs)*time.Second)
	defer cancel()

	dataM, errConv := jsonToBsonM(data)
	if errConv != nil {
		return primitive.ObjectID{}, errConv
	}

	m.l.Debug("Bison format data to insert: ", dataM)

	collection := m.client.Database(m.Database).Collection(m.Collection)
	if collection == nil {
		return primitive.ObjectID{}, errors.New("collection is nil")
	}

	m.l.Debug("collection returned: ", collection)

	result, errInsert := collection.InsertOne(ctxLocal, dataM)
	m.l.Debug("result: ", result)

	if result == nil {
		return primitive.ObjectID{}, errInsert
	}

	return result.InsertedID.(primitive.ObjectID), errInsert
}

// FindOne Method finds data based on passed filter and returns it.
func (m *MoInstance) FindOne(ctx context.Context, filter []byte) (interface{}, error) {
	ctxLocal, cancel := context.WithTimeout(ctx, time.Duration(m.TimeoutSecs)*time.Second)
	defer cancel()

	bsonFilter, errConv := jsonToBsonM(filter)
	if errConv != nil {
		return nil, errConv
	}

	var result bson.M
	if errFind := m.client.Database(m.Database).Collection(m.Collection).FindOne(ctxLocal, bsonFilter).Decode(&result); errFind != nil {
		return nil, errFind
	}
	return result, nil
}

// FindByID Method finds data based on passed ID and returns it.
func (m *MoInstance) FindByID(ctx context.Context, id primitive.ObjectID) (interface{}, error) {
	ctxLocal, cancel := context.WithTimeout(ctx, time.Duration(m.Cfg.TimeoutSecs)*time.Second)
	defer cancel()

	bsonFilter := bson.M{"_id": bson.M{"$eq": id}}

	m.l.Debug("Bison format data to find: ", bsonFilter)

	var result bson.M
	errFind := m.client.Database(m.Database).Collection(m.Cfg.Collection).FindOne(ctxLocal, bsonFilter).Decode(&result)
	if errFind != nil {
		return nil, errFind
	}
	return result, nil // ex. "5d678d799139918d230cfd41"
}

// FindManyFilterJSON Method finds data based on passed ID and returns it. Could return more than one record.
func (m *MoInstance) FindManyFilterJSON(ctx context.Context, filterJSON []byte) ([]bson.M, error) {
	ctxLocal, cancel := context.WithTimeout(ctx, time.Duration(m.Cfg.TimeoutSecs)*time.Second)
	defer cancel()

	bsonFilter, errConv := jsonToBsonM(filterJSON)
	if errConv != nil {
		return nil, errConv
	}

	cursor, errFind := m.client.Database(m.Cfg.Database).Collection(m.Cfg.Collection).Find(ctxLocal, bsonFilter)
	if errFind != nil {
		return nil, errFind
	}
	defer cursor.Close(ctxLocal)

	return walkMongoSet(ctxLocal, cursor)
}

// FindManyFilterBSON Method finds data based on passed ID and returns it. Could return more than one record.
func (m *MoInstance) FindManyFilterBSON(ctx context.Context, filterBSON primitive.M) ([]bson.M, error) {
	ctxLocal, cancel := context.WithTimeout(ctx, time.Duration(m.Cfg.TimeoutSecs)*time.Second)
	defer cancel()

	cursor, errFind := m.client.Database(m.Cfg.Database).Collection(m.Cfg.Collection).Find(ctxLocal, filterBSON)
	if errFind != nil {
		return nil, errFind
	}
	defer cursor.Close(ctxLocal)

	return walkMongoSet(ctxLocal, cursor)
}

func walkMongoSet(ctx context.Context, cursor *mongo.Cursor) ([]bson.M, error) {
	var result []bson.M

	for cursor.Next(ctx) {
		var buf bson.M

		if errDecode := cursor.Decode(&buf); errDecode != nil {
			return nil, errors.Wrap(errDecode, "could not decode into buffer")
		}
		result = append(result, buf)
	}

	if errCursor := cursor.Err(); errCursor != nil {
		return nil, errors.Wrap(errCursor, "cursor error")
	}
	return result, nil
}

// DeleteOne Method deletes one record from found.
func (m *MoInstance) DeleteOne(ctx context.Context, filter []byte) (interface{}, error) {
	ctxLocal, cancel := context.WithTimeout(ctx, time.Duration(m.Cfg.TimeoutSecs)*time.Second)
	defer cancel()

	bsonFilter, errConv := jsonToBsonM(filter)
	if errConv != nil {
		return nil, errConv
	}

	return m.client.Database(m.Cfg.Database).Collection(m.Cfg.Collection).DeleteOne(ctxLocal, bsonFilter)
}

// DeleteAll Method deletes all records found matching passed filter.
func (m *MoInstance) DeleteAll(ctx context.Context, filter []byte) (interface{}, error) {
	ctxLocal, cancel := context.WithTimeout(ctx, time.Duration(m.TimeoutSecs)*time.Second)
	defer cancel()

	bsonFilter, errConv := jsonToBsonM(filter)
	if errConv != nil {
		return nil, errConv
	}

	return m.client.Database(m.Cfg.Database).Collection(m.Cfg.Collection).DeleteMany(ctxLocal, bsonFilter)
}

// UpdateByID Method updates record with passed ID.
func (m *MoInstance) UpdateByID(ctx context.Context, id primitive.ObjectID, newValue bson.M) (interface{}, error) {
	ctxLocal, cancel := context.WithTimeout(ctx, time.Duration(m.TimeoutSecs)*time.Second)
	defer cancel()

	bsonFilter := bson.M{"_id": bson.M{"$eq": id}}

	return m.client.Database(m.Database).Collection(m.Collection).UpdateOne(ctxLocal, bsonFilter, newValue)
}

// UpdateOne Method updates one record from those matching passed filter.
func (m *MoInstance) UpdateOne(ctx context.Context, filter primitive.M, newValue bson.M) (interface{}, error) {
	ctxLocal, cancel := context.WithTimeout(ctx, time.Duration(m.TimeoutSecs)*time.Second)
	defer cancel()

	return m.client.Database(m.Database).Collection(m.Collection).UpdateOne(ctxLocal, filter, newValue)
}

// UpdateMany Method updates all records that match the passed filter search.
func (m *MoInstance) UpdateMany(ctx context.Context, filter []byte, newValue bson.M) (interface{}, error) {
	ctxLocal, cancel := context.WithTimeout(ctx, time.Duration(m.TimeoutSecs)*time.Second)
	defer cancel()

	bsonFilter, errConv := jsonToBsonM(filter)
	if errConv != nil {
		return nil, errConv
	}

	return m.client.Database(m.Database).Collection(m.Collection).UpdateMany(ctxLocal, bsonFilter, newValue)
}
