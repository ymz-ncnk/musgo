package musgo

import "errors"

// // ErrNotAlias happens on building alias TypeDesc for not alias type.
var ErrNotAlias = errors.New("not an alias")