package main

import (
	"fmt"
)

const src = "
package something

import (
	\"fmt\"
)

// Documentation for SomeType
// STRUCT START OMIT
type SomeType struct {
	FieldA int `json:\"field_a\"`
	fmt.Stringer
}
// STRUCT END OMIT

// Documentation for MethodA
func (st SomeType) MethodA() {
	fmt.Println(\"Method A\")
}

func (st SomeType) methodB() {
	fmt.Println(\"Method B\")
}
"
