package main

import (
	"log"
	"path"
	"runtime"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Dir(filename)

	if err := entc.Generate(dir+"/../ent/schema/", &gen.Config{
		Schema:   "github.com/ent/ent/entc/integration/ulid/ent",
		Target:   dir + "/../ent",
		Package:  "github.com/ent/ent/entc/integration/ulid/ent",
		Features: []gen.Feature{},
	}); err != nil {
		log.Fatalf("error running ent codegen: %v", err)
	}
}
