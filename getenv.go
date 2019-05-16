package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println(filepath.Clean(runtime.GOROOT()))
}
