package slices

func Swap[T any](s []T, idxA, idxB int) {
	s[idxB], s[idxA] = s[idxA], s[idxB]
}

func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Prepend[T any](s *[]T, x T) {
	*s = append([]T{x}, (*s)...)
}
