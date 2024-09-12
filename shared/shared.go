package shared

import "fmt"

func Pointer[T any](x T) *T {
	return &x
}

func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

func DigAny(source any, key string) any {
	return source.(map[string]any)[key]
}

func Dig[T any](source any, key string) T {
	v, found := source.(map[string]any)[key]
	if !found {
		return Default[T]()
	}
	v1, isT := v.(T)
	if !isT {
		panic(fmt.Sprintf("dug value for key %s not convertible: %#v", key, v))
	}
	return v1
}

func DigDelete[T any](source any, key string) T {
	source1 := source.(map[string]any)
	v, found := source1[key]
	if !found {
		return Default[T]()
	}
	delete(source1, key)
	v1, isT := v.(T)
	if !isT {
		panic(fmt.Sprintf("dug value for key %s not convertible: %#v", key, v))
	}
	return v1
}

func Default[T any]() T {
	var out T
	return out
}

func AnyIsNotZero(v any, fallback bool) bool {
	switch v1 := v.(type) {
	case bool:
		return v1
	case string:
		return v1 != ""
	case int, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return v1 != 0
	case float32, float64:
		return v1 != 0.0
	}
	return fallback
}

func InMap(key string, m map[string]any) bool {
	v, res := m[key]
	if !res {
		return false
	}
	return AnyIsNotZero(v, res)
}
