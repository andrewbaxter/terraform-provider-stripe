package shared

func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}
