package scalar

import (
	"encoding/hex"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

func UnmarshalKID(v interface{}) (string, error) {
	kid, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("expect type 'string' and got type '%T'", v)
	}

	raw, err := hex.DecodeString(kid)
	if err != nil {
		return "", err
	}
	if len(raw) != 35 {
		return "", fmt.Errorf("expect a 35 bytes hex encoded key")
	}

	return kid, nil
}

func MarshalKID(kid string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(`"`))
		_, _ = w.Write([]byte(kid))
		_, _ = w.Write([]byte(`"`))
	})
}
