package musgo

// Conf configures the generation process.
type Conf struct {
	Unsafe bool   // Generate unsafe code or not.
	Suffix string // Suffix for Marshal, Unmarshal, Size methods.
	Path   string // Folder of the generated file.
}

// AliasConf configures the generation process for an alias type.
type AliasConf struct {
	Conf
	Validator string // Validates value.
	Encoding  string // Sets encoding.
	MaxLength int    // If alias to string, array, slice, or map, restricts
	// length, should be positive number.
	ElemValidator string // If alias to array, slice, or map, validates elements.
	ElemEncoding  string // If alias to array, slic, or map, sets encoding.
	KeyValidator  string // If alias to map, validates keys.
	KeyEncoding   string // If alias to map, sets encoding.
}
