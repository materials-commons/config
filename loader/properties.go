package loader

import (
	p "github.com/dmotylev/goproperties"
	"github.com/materials-commons/config"
	"io"
)

type propertiesLoader struct {
	r io.Reader
}

func Properties(r io.Reader) config.Loader {
	return &propertiesLoader{r: r}
}

func (l *propertiesLoader) Load(out interface{}) error {
	var properties p.Properties
	if err := properties.Load(l.r); err != nil {
		return err
	}
	out = properties
	return nil
}
