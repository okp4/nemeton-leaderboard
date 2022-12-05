package scalar

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/pkg/errors"
)

func UnmarshalJSON(v interface{}) (map[string]interface{}, error) {
	switch val := v.(type) {
	case map[string]interface{}:
		return val, nil
	case string:
		var jsonMap map[string]interface{}
		if err := json.Unmarshal([]byte(val), &jsonMap); err != nil {
			return nil, errors.Wrap(err, "couldn't parse json")
		}
		return jsonMap, nil
	default:
		return nil, fmt.Errorf("unsupported type")
	}
}

func MarshalJSON(jsonMap map[string]interface{}) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		raw, _ := json.Marshal(jsonMap)
		_, _ = w.Write(raw)
	})
}
