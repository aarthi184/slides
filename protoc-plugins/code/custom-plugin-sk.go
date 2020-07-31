package main

import (
	"plugin/custom"

	"github.com/gogo/protobuf/vanity/command"
)

func main() {
	req := command.Read()

	p := custom.New()

	resp := command.GeneratePlugin(req, p, "_custom.gen.go")

	command.Write(resp)
}
