package custom

import "github.com/gogo/protobuf/protoc-gen-gogo/generator"

type custom struct {
	*generator.Generator
}

func (p *custom) Generate(file *generator.FileDescriptor) {
	for _, service := range file.Service {
		for _, method := range service.Method {

			p.P("func (c *client) methodName(inputParams) (outputParams) {")
			p.In()
			...
			p.Out()
			p.P("}")
		}
	}
}

