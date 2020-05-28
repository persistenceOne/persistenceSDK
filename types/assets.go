package types

type Assets interface {
	String() string

	ID() ID
	Asset(ID) Asset

	Add(Asset) error
	Remove(Asset) error
	Mutate(Asset) error
}