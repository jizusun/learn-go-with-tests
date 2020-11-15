package di

import (
	"fmt"
	"io"
)

// Greet to write to console
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// func main() {
// 	Greet("jizu")
// }
