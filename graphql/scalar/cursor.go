package scalar

import (
	"fmt"
	"io"

	"okp4/nemeton-leaderboard/app/nemeton"

	"github.com/99designs/gqlgen/graphql"
)

func UnmarshalCursor(v interface{}) (*nemeton.Cursor, error) {
	strCursor, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("expect type 'string' and got type '%T'", v)
	}

	c := &nemeton.Cursor{}
	return c, c.FromHex(strCursor)
}

func MarshalCursor(c *nemeton.Cursor) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(`"`))
		_, _ = w.Write([]byte(c.Hex()))
		_, _ = w.Write([]byte(`"`))
	})
}
