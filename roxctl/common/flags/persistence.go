package flags

import (
	"fmt"

	"github.com/stackrox/rox/pkg/renderer"
)

// PersistenceTypeWrapper implements the pflags.Value interface for persistence type.
type PersistenceTypeWrapper struct {
	PersistenceType *renderer.PersistenceType
}

// String implements the pflags.Value interface.
func (f *PersistenceTypeWrapper) String() string {
	return f.PersistenceType.String()
}

// Set implements the pflags.Value interface.
func (f *PersistenceTypeWrapper) Set(input string) error {
	pt, ok := renderer.StringToPersistentTypes[input]
	if !ok {
		return fmt.Errorf("Invalid persistence type: %s", input)
	}
	*f.PersistenceType = pt
	return nil
}

// Type implements the pflags.Value interface.
func (f *PersistenceTypeWrapper) Type() string {
	return "persistence type"
}
