package mocks

import (
	"github.com/ymz-ncnk/amock"
	"github.com/ymz-ncnk/musgen/v2"
)

func NewPersistor() Persistor {
	return Persistor{amock.New("Persistor")}
}

type Persistor struct {
	*amock.Mock
}

func (persistor Persistor) RegisterPersist(
	fn func(tDesc musgen.TypeDesc, data []byte, path string) error) Persistor {
	persistor.Register("Persist", fn)
	return persistor
}

func (persistor Persistor) Persist(tDesc musgen.TypeDesc, data []byte,
	path string) (err error) {
	vals, err := persistor.Call("Persist", tDesc, data, path)
	if err != nil {
		return
	}
	err, _ = vals[0].(error)
	return
}
