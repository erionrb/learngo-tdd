package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}


func main() {
	Greet(os.Stdout, "Erion")
}