// executor/runtime.go

package runtime

type Runtime struct {
	store map[string]interface{}
}

func NewRuntime() *Runtime {
	return &Runtime{
		store: make(map[string]interface{}),
	}
}

func (r *Runtime) Get(name string) interface{} {
	if value, ok := r.store[name]; ok {
		return value
	}
	// Handle variable not found, return nil or throw error
	return nil
}

func (r *Runtime) Set(name string, value interface{}) {
	r.store[name] = value
}
