// +build ignore
package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

const src = `
// CODE START OMIT
package something

import (
	"fmt"
)

// Documentation for SomeType
// STRUCT START OMIT
type SomeType struct {
	FieldA int "json:\"field_a\""
	fmt.Stringer
}
// STRUCT END OMIT

// Documentation for MethodA
func (st SomeType) MethodA() {
	fmt.Println("Method A")
}

func (st SomeType) methodB() {
	fmt.Println("Method B")
}
// CODE END OMIT
`

func main() {
	// PARSE START OMIT
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "src.go", src, 0)
	ast.Print(fset, f)
	// PARSE END OMIT
}
