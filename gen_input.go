package main

// GEN SIMPLE START OMIT
// +gen
type SomeType struct {
	//...
}

// GEN SIMPLE END OMIT

// GEN POINTER START OMIT
// +gen *
type AnotherType struct {
	//...
}

// GEN POINTER END OMIT

// GEN METHODS START OMIT
// +gen methods:"Any,Where,Count"
type YetAnotherType struct {
	//...
}

// GEN METHODS END OMIT

// GEN PROJECTIONS START OMIT
// +gen projections:"int,string"
type DifferentType struct {
	//...
}

// GEN PROJECTIONS END OMIT

// GEN CONTAINERS START OMIT
// +gen containers:"List,Set,Ring"
type AnotherDifferentType struct {
	//...
}

// GEN CONTAINERS END OMIT
