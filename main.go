package main

import (
	"fmt"

	"example.com/go_package_deps/a"
	"example.com/go_package_deps/b"
)

func main() {
	a.A()
	fmt.Println(b.B)
}
