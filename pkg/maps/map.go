package maps

func ComputeIfAbsent[K comparable, V any](collection map[K]V, key K, fc func(K) V) V {
	v, exist := collection[key]
	if exist {
		return v
	}
	v = fc(key)
	collection[key] = v
	return v
}
