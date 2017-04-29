package demo

// Tomate is about red vegetables to make famous italian food.
type Tomate struct {
	Name string
}

// GetID return the ID of the Tomate.
func (t Tomate) GetID() string {
	return t.Name
}

//go:generate lister vegetables_gen.go Tomate:Tomates
//go:generate mutexer vegetuxed_gen.go Tomates:Tomatex
