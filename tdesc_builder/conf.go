package tdesc_builder

// Conf configuration for a struct or alias type.
type Conf struct {
	Unsafe bool   // generate unsafe code or not
	Suffix string // suffix for Marshal, Unmarshal, Size methods
}

// -----------------------------------------------------------------------------
// AliasConf configuration with validators and encodings for an alias type.
type AliasConf struct {
	Conf
	Validator string // validates value
	Encoding  string // sets encoding
	MaxLength int    // if alias to string, array, slice, or map, restricts
	// length, should be positive number
	ElemValidator string // if alias to array, slice, or map, validates elements
	ElemEncoding  string // if alias to array, slic, or map, sets encoding
	KeyValidator  string // if alias to map, validates keys
	KeyEncoding   string // if alias to map, sets encoding
}
