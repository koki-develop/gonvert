package util

func Contains[T comparable](ts []T, t T) bool {
	for _, l := range ts {
		if l == t {
			return true
		}
	}
	return false
}
