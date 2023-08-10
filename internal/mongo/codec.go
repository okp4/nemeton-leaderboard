package mongo

import (
	"reflect"

	"github.com/cosmos/cosmos-sdk/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
)

var (
	tAccAddress  = reflect.TypeOf(types.AccAddress{})
	tValAddress  = reflect.TypeOf(types.ValAddress{})
	tConsAddress = reflect.TypeOf(types.ConsAddress{})
)

func MakeRegistry() *bsoncodec.Registry {
	return bson.NewRegistryBuilder().
		RegisterTypeEncoder(tAccAddress, bsoncodec.ValueEncoderFunc(encodeAccAddress)).
		RegisterTypeEncoder(tValAddress, bsoncodec.ValueEncoderFunc(encodeValAddress)).
		RegisterTypeEncoder(tConsAddress, bsoncodec.ValueEncoderFunc(encodeConsAddress)).
		Build()
}

func encodeAccAddress(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tAccAddress {
		return bsoncodec.ValueEncoderError{Name: "ObjectIDEncodeValue", Types: []reflect.Type{tAccAddress}, Received: val}
	}

	return vw.WriteString(val.Interface().(types.AccAddress).String())
}

func encodeValAddress(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tValAddress {
		return bsoncodec.ValueEncoderError{Name: "ObjectIDEncodeValue", Types: []reflect.Type{tValAddress}, Received: val}
	}

	return vw.WriteString(val.Interface().(types.ValAddress).String())
}

func encodeConsAddress(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tConsAddress {
		return bsoncodec.ValueEncoderError{Name: "ObjectIDEncodeValue", Types: []reflect.Type{tConsAddress}, Received: val}
	}

	return vw.WriteString(val.Interface().(types.ConsAddress).String())
}
