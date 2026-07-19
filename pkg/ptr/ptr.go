package ptr

// Or returns the value p points at, or fallback when p is nil.
// Partial-update inputs use nil to mean "leave this field alone".
func Or[T any](p *T, fallback T) T {
	if p == nil {
		return fallback
	}
	return *p
}
