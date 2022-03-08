package codec

import (
"go.mongodb.org/mongo-driver/bson"
)

type BSONCodec struct {}

func (c BSONCodec) Marshal(v interface{}) ([]byte, error) {
	return bson.Marshal(v)
}

func (c BSONCodec) Unmarshal(data []byte, v interface{}) error {
	return bson.Unmarshal(data, v)
}

func (c BSONCodec) Name() string {
	return "bsonCodec"
}