package shared

import "sync"

func Memoize[K comparable, V any](fn func(K) (V, error)) func(K) (V, error) {
	var (
		mu sync.RWMutex
		c  = make(map[K]V)
	)
	return func(key K) (V, error) {
		mu.RLock()
		v, ok := c[key]
		mu.RUnlock()
		if ok {
			return v, nil
		}

		mu.Lock()
		defer mu.Unlock()

		if v, ok = c[key]; ok {
			return v, nil
		}

		res, err := fn(key)
		if err != nil {
			var zero V
			return zero, err
		}
		c[key] = res
		return res, nil
	}
}
