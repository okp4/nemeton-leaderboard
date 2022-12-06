package nemeton

// BlockRange Represents a blockchain block range.
type BlockRange struct {
	// The block height the range begin, inclusive.
	From int `bson:"from"`
	// The block height the range end, inclusive.
	To int `bson:"to"`
}
