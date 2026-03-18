package ptr

func Or[T any](input *T, fallback T) T {
	if input != nil {
		return *input
	}

	return fallback
}

func To[T any](v T) *T {
	return &v
}
