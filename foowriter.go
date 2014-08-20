package foowriter

import (
	"fmt"
	"io"

	"github.com/clipperhouse/gen/typewriter"
)

func init() {
	// INIT START OMIT
	err := typewriter.Register(&fooWriter{})
	// INIT END OMIT
	if err != nil {
		panic(err)
	}
}

type fooWriter struct{}

func (f *fooWriter) Name() string {
	return "foo"
}

func (f *fooWriter) Validate(t typewriter.Type) (bool, error) {
	return true, nil
}

func (f *fooWriter) WriteHeader(w io.Writer, t typewriter.Type) {
	w.Write([]byte("// foowriter was here"))
	return
}

func (f *fooWriter) Imports(t typewriter.Type) (result []typewriter.ImportSpec) {
	result = append(result, typewriter.ImportSpec{Path: "fmt"})
	return result
}

func (f *fooWriter) WriteBody(w io.Writer, t typewriter.Type) {
	w.Write([]byte(fmt.Sprintf(`func pointless%s(){
		fmt.Println("pointless!")
		}`, t.LocalName())))
	return
}
