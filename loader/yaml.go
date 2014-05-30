package loader

import (
	"bytes"
	"github.com/materials-commons/config"
	"gopkg.in/yaml.v1"
	"io"
)

type yamlLoader struct {
	r io.Reader
}

func YAML(r io.Reader) config.Loader {
	return &yamlLoader{r: r}
}

func (l *yamlLoader) Load(out interface{}) error {
	var buf bytes.Buffer
	buf.ReadFrom(l.r)
	if err := yaml.Unmarshal(buf.Bytes(), out); err != nil {
		return err
	}
	return nil
}
