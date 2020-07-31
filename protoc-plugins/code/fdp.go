
type FileDescriptorProto struct {
	Name       *string
	Package    *string
	Dependency []string

	...

	// All top-level definitions in this file.
	MessageType []*DescriptorProto
	EnumType    []*EnumDescriptorProto
	Service     []*ServiceDescriptorProto
	Extension   []*FieldDescriptorProto
	Options     *FileOptions

	...

	Syntax *string
}
