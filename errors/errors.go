package errors

type Errors map[string][]string

func NewErrors() Errors {
	return make(Errors)
}

// 增加 error
func (e Errors) AddError(key, err string) {
	e[key] = append(e[key], err)
}

// 删除 error
func (e Errors) Clear() {
	keys := make([]string, 0, len(e))
	for k := range e {
		keys = append(keys, k)
	}

	for _, key := range keys {
		delete(e, key)
	}
}
