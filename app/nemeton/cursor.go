package nemeton

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cursor struct {
	points   uint64
	objectID primitive.ObjectID
}

func (c *Cursor) Hex() string {
	return hex.EncodeToString(c.MarshalBinary())
}

func (c *Cursor) FromHex(enc string) error {
	bin, err := hex.DecodeString(enc)
	if err != nil {
		return err
	}

	return c.UnmarshalBinary(bin)
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
