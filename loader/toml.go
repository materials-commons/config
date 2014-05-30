package loader

import (
	"github.com/BurntSushi/toml"
	"github.com/materials-commons/config"
	"io"
)

type tomlLoader struct {
	r io.Reader
}

func TOML(r io.Reader) config.Loader {
	return &tomlLoader{r: r}
}

func (l *tomlLoader) Load(out interface{}) error {
	if _, err := toml.DecodeReader(l.r, out); err != nil {
		return err
	}
	return nil
}
