package offset

type Offset interface {
	Marshal() string
	Unmarshal(from string)
}
