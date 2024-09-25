package from

func NoValue[V any](int, V) struct{} { return struct{}{} }

func Slice[K comparable, V1 any, V2 any, S ~[]V1](s S, key func(int, V1) K, value func(int, V1) V2) map[K]V2 {
	m := map[K]V2{}
	for i, v := range s {
		m[key(i, v)] = value(i, v)
	}
	return m
}
