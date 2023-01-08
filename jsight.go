package main

import (
	"fmt"
	"plugin"
)

func loadJSightPlugin(path string) {
	p, err := plugin.Open(path)
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("ValidationError")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("JSightValidateHTTPRequest")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v, %#v", f, v)
}
