package config

type Handler interface {
	Load(path string) error
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}) error
}

type jsonHandler struct {
	values map[string]interface{}
}

func JSONHandler() Handler {
	return &jsonHandler{
		values: make(map[string]interface{}),
	}
}

func (h *jsonHandler) Load(path string) error {
	return nil
}

func (h *jsonHandler) Get(key string) (interface{}, bool) {
	val, found := h.values[key]
	return val, found
}

func (h *jsonHandler) Set(key string, value interface{}) error {
	return nil
}




















