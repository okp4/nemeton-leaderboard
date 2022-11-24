package scalar

import (
	"bytes"
	"fmt"
	"io"
	"net/url"

	"github.com/99designs/gqlgen/graphql"
)

func UnmarshalURI(v interface{}) (*url.URL, error) {
	strURL, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("expect type 'string' and got type '%T'", v)
	}

	return url.Parse(strURL)
}

func MarshalURI(uri *url.URL) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		var buffer bytes.Buffer

		buffer.WriteRune('"')
		buffer.WriteString(uri.String())
		buffer.WriteRune('"')

		_, _ = buffer.WriteTo(w)
	})
}
