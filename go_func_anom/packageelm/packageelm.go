package main

import (
	"fmt"
	"go/importer"
	"go/types"
)

func main() {
	pkg, err := importer.Default().Import("time")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	scope := pkg.Scope()
	for _, name := range scope.Names() {
		if name == "Time" {
			obj := scope.Lookup(name)
			if tn, ok := obj.Type().(*types.Named); ok {
				fmt.Printf("%#v\n", tn.NumMethods())
			}
		}
	}
}
