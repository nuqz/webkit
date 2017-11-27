package db

// Model is the interface that represents one record in the database.
type Model interface {
	// AsMap converts model to map[string]interface{}.
	AsMap() map[string]interface{}

	// IsNew returns true if model is not stored in the database, otherwise it
	// returns false.
	IsNew() bool

	// PK returns primary key column (field, attribute, etc...) name as string
	// and it's value as empty interface.
	PK() (string, interface{})

	// Store used to set "is new" flag to false when saving a model to the
	// database.
	Store()
}

// M implements IsNew and Store methods and used for embedding it into actual
// models.
type M struct {
	isNew bool // End user should not be able to set it manually.
}

// NewM returns new M and sets isNew to true.
func NewM() *M { return &M{true} }

// IsNew implements Model interface.
func (m *M) IsNew() bool { return m.isNew }

// Store implements Model interface.
func (m *M) Store() { m.isNew = false }

// NewAndExisting divides the slice of models ms into two new slices. The first
// one contains new models and the second one contains models, which are present
// in the database.
func NewAndExisting(ms []Model) ([]Model, []Model) {
	newMs, oldMs := make([]Model, 0, 0), make([]Model, 0, 0)

	for _, m := range ms {
		if m.IsNew() {
			newMs = append(newMs, m)
			continue
		}

		oldMs = append(oldMs, m)
	}

	return newMs, oldMs
}

// PKs returns a new slice with the primary key values of models in ms slice.
func PKs(ms []Model) []interface{} {
	ln := len(ms)
	keys := make([]interface{}, ln, ln)

	for i := 0; i < ln; i++ {
		if _, val := ms[i].PK(); !ms[i].IsNew() {
			keys[i] = val
		}
	}

	return keys
}
