package nemeton

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/btcutil/base58"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cursor struct {
	points   uint64
	objectID primitive.ObjectID
}

func (c *Cursor) Marshal() string {
	return base58.Encode(c.MarshalBinary())
}

func (c *Cursor) Unmarshal(enc string) error {
	return c.UnmarshalBinary(base58.Decode(enc))
}

func (c *Cursor) MarshalBinary() []byte {
	var enc [20]byte

	binary.LittleEndian.PutUint64(enc[:8], c.points)

	rawID := [12]byte(c.objectID)
	copy(enc[8:], rawID[:])

	return enc[:]
}

func (c *Cursor) UnmarshalBinary(enc []byte) error {
	if len(enc) != 20 {
		return fmt.Errorf("expected a 20 length byte array")
	}

	c.points = binary.LittleEndian.Uint64(enc[:8])

	var decID [12]byte
	copy(decID[:], enc[8:])
	c.objectID = decID

	return nil
}
