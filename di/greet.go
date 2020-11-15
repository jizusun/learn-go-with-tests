package di

import (
	"fmt"
	"io"
)

// Greet write to io.Writer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
