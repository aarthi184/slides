
type Plugin interface {

	// Name identifies the plugin.
	Name() string

	// Init is called once after data structures are built but before
	// code generation begins.
	Init(g *Generator)

	// Generate produces the code generated by the plugin for this file,
	// except for the imports, by calling the generator's methods P, In, and Out.
	Generate(file *FileDescriptor)

	// GenerateImports produces the import declarations for this file.
	// It is called after Generate.
	GenerateImports(file *FileDescriptor)


}

