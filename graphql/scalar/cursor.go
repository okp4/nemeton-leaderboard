package scalar

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UnmarshalCursor(v interface{}) (primitive.ObjectID, error) {
	strCursor, ok := v.(string)
	if !ok {
		return primitive.ObjectID{}, fmt.Errorf("expect type 'string' and got type '%T'", v)
	}

	return primitive.ObjectIDFromHex(strCursor)
}

func MarshalCursor(c primitive.ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(c.Hex()))
	})
}
