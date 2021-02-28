package injection

import (
	"fmt"
	"io"
)

func Greet(w io.Writer, msg string) {
	fmt.Fprintf(w, "Hello, %s", msg)
}
