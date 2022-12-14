package nemeton

// BlockRange Represents a blockchain block range.
type BlockRange struct {
	// The block height the range begin, inclusive.
	From int64 `bson:"from"`
	// The block height the range end, inclusive.
	To int64 `bson:"to"`
}
