package cache

var cache map[string]any

func InitializeCache() {
	cache = make(map[string]any)
}

func Get(key string) (any, error) {
	return cache[key], nil
}

func Set(key string, value any) error {
	cache[key] = value
	return nil
}
