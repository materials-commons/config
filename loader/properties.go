package loader

import (
	p "github.com/dmotylev/goproperties"
	"github.com/materials-commons/config"
	"io"
)

type propertiesLoader struct {
	r io.Reader
}

// Properties creates a new Loader for properties formatted data.
func Properties(r io.Reader) config.Loader {
	return &propertiesLoader{r: r}
}

// Load loads the data from the reader.
func (l *propertiesLoader) Load(out interface{}) error {
	var properties p.Properties
	if err := properties.Load(l.r); err != nil {
		return err
	}
	out = properties
	return nil
}
