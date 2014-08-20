// +build ignore
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"
)

type VisitFunc func(node ast.Node) ast.Visitor

func (vf VisitFunc) Visit(node ast.Node) ast.Visitor {
	return vf(node)
}

const src = `
package main

import (
	"fmt"
)

func add2Int(i int) int {
	return i + 2
}

func add2Float64(i float64) float64 {
	return i + 2.0
}

func main() {
	// GENERIC CODE START OMIT
	seriesA := []int{0, 1, 2, 3}
	seriesA = _Map_Ints(seriesA, add2Int)

	seriesB := []float64{0, 1, 2, 3}
	seriesB = _Map_Float64s(seriesB, add2Float64)
	// GENERIC CODE END OMIT
}`

// PARSE START OMIT
func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "src.go", src, 0)
	ast.Walk(VisitFunc(PrintVisitor), f)
}

// PARSE END OMIT

var lastNode ast.Node

// VISITOR START OMIT
func PrintVisitor(node ast.Node) (w ast.Visitor) {
	w = VisitFunc(PrintVisitor)
	switch x := node.(type) {
	case *ast.CallExpr:
		// VISITOR END OMIT
		f, ok := x.Fun.(*ast.Ident)
		// Not a stand-alone function call
		if !ok {
			return
		}

		if !strings.HasPrefix(f.Name, "_") {
			return
		}

		elems := strings.Split(f.Name, "_")
		if len(elems) < 2 {
			log.Fatalf("Invalid template syntax: %s", f.Name)
		}

		stpl, ok := templates[elems[1]]
		if !ok {
			log.Fatalf("Invalid template name: %s", elems[1])
		}

		tpl := template.New("")
		tpl.Funcs(template.FuncMap{
			"typestring": typestring,
			"elem":       elem,
		})
		tpl = template.Must(tpl.Parse(stpl))

		// TEMPLATE RENDERING START OMIT
		tpl.Execute(os.Stdout, map[string]interface{}{
			"Name": x.Fun.(*ast.Ident).Name,
			"Args": x.Args,
		})
		// TEMPLATE RENDERING END OMIT
	}
	return
}

// TEMPLATE START OMIT
var templates = map[string]string{
	"Map": `func {{.Name}}(s {{index .Args 0 | typestring}}, f func({{index .Args 0 | typestring | elem}}) {{index .Args 0 | typestring | elem}}) {{index .Args 0 | typestring}} {
	r := make({{index .Args 0 | typestring}}, 0, len(s))
	for i := range s {
		r = append(r, f(s[i]))
	}
	return r
}
`,
}

// TEMPLATE END OMIT

func elem(x string) string {
	return strings.TrimPrefix(x, "[]")
}

func typestring(x interface{}) string {
	switch x := x.(type) {
	case *ast.AssignStmt:
		return typestring(x.Rhs[0])
	case *ast.CompositeLit:
		return typestring(x.Type)
	case *ast.ArrayType:
		return "[]" + typestring(x.Elt)
	case *ast.Ident:
		if x.Obj == nil {
			return x.Name
		}
		return typestring(x.Obj.Decl)
	default:
		return fmt.Sprintf("%#v", x)
	}
}
