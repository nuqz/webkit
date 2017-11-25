package db

// Storable is the interface that wraps IsNew and Store methods.
//
// IsNew returns true if model is not stored in the database, otherwise it
// returns false.
//
// Store used to set "is new" flag to false when saving a model to the database.
type Storable interface {
	IsNew() bool
	Store()
}

// S is the implementation of IsNewer and Storer interfaces.
type S struct {
	isNew bool // End user should not be able to set it manually.
}

// NewS returns new S and sets isNew to true.
func NewS() *S { return &S{true} }

// IsNew implements Storable interface.
func (s *S) IsNew() bool { return s.isNew }

// Store implements Storable interface.
func (s *S) Store() { s.isNew = false }

// PKer is the interface that wraps PK method.
// PK returns primary key column (field, attribute, etc...) name as string and
// it's value as empty interface.
type PKer interface {
	PK() (string, interface{})
}

// Model is the interface that wraps IsNewer and PKer interfaces.
// Almost in every case we should know whether the model is new and also to know
// primary key's name and value.
type Model interface {
	Storable
	PKer
}
