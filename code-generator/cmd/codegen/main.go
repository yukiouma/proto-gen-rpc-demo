package main

import (
	"codegen/internal/codegen"
	"log"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	g, err := codegen.NewGenerator(codegen.GoCodeTemplate)
	if err != nil {
		log.Fatalf("[ERROR]: failed to initialize code generator")
	}
	protogen.Options{}.Run(g.Generate)
}
