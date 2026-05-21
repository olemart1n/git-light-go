package object

type GitObject interface {
	Type() string
	// Hash() string
	Body() []byte
}
