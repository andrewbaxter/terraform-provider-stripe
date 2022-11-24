package shared

func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

func Dig[T any](bla any, field string) T {
	v, found := bla.(map[string]any)[field]
	if !found {
		return Default[T]()
	}
	v1, isT := v.(T)
	if !isT {
		panic("dug value not convertible")
	}
	return v1
}

func DigDelete[T any](bla any, field string) T {
	bla1 := bla.(map[string]any)
	v, found := bla1[field]
	if !found {
		return Default[T]()
	}
	delete(bla1, field)
	v1, isT := v.(T)
	if !isT {
		panic("dug value not convertible")
	}
	return v1
}

func Default[T any]() T {
	var out T
	return out
}
