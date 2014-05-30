package loader

import (
	"bytes"
	"encoding/json"
	"github.com/materials-commons/config"
	"io"
)

type jsonLoader struct {
	r io.Reader
}

func JSON(r io.Reader) config.Loader {
	return &jsonLoader{r: r}
}

func (l *jsonLoader) Load(out interface{}) error {
	var buf bytes.Buffer
	buf.ReadFrom(l.r)
	if err := json.Unmarshal(buf.Bytes(), out); err != nil {
		return err
	}
	return nil
}
