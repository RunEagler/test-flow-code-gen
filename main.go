package main

import (
	"flag"
	"test-flow-code-gen/testflow"
)

func main() {
	targetDir := flag.String("path", "examples/services", "./services")
	flag.Parse()
	testflow.GenerateTestTemplate(*targetDir)
}
