package map2d

import (
	"cmp"
)

type Map2D[K comparable, T any] map[K]map[K]T

func New[K comparable, T any](cap int) Map2D[K, T] {
	return make(Map2D[K, T], cap)
}

func (m *Map2D[K, T]) Set(k1, k2 K, x T) *Map2D[K, T] {
	if _, ok := (*m)[k1]; !ok {
		(*m)[k1] = make(map[K]T)
	}

	(*m)[k1][k2] = x

	return m
}

func (m *Map2D[K, T]) Get(k1, k2 K) (T, bool) {
	if _, ok := (*m)[k1]; !ok {
		return *new(T), false
	}

	return (*m)[k1][k2], true
}

type Dp[K comparable, T cmp.Ordered] map[K]map[K]T

func NewDp[K comparable, T cmp.Ordered](cap int) Dp[K, T] {
	return make(Dp[K, T], cap)
}

func (m *Dp[K, T]) Set(k1, k2 K, x T) *Dp[K, T] {
	if _, ok := (*m)[k1]; !ok {
		(*m)[k1] = make(map[K]T)
	}

	(*m)[k1][k2] = x

	return m
}

func (m *Dp[K, T]) Del(k1 K) *Dp[K, T] {
	delete(*m, k1)
	return m
}

func (m *Dp[K, T]) SubMap(k1 K) (map[K]T, bool) {
	if sub, ok := (*m)[k1]; !ok {
		return nil, false
	} else {
		return sub, true
	}
}

func (m *Dp[K, T]) Get(k1, k2 K) (T, bool) {
	if _, ok := (*m)[k1]; !ok {
		return *new(T), false
	}

	return (*m)[k1][k2], true
}

func (m *Dp[K, T]) SetMin(k1, k2 K, x T) *Dp[K, T] {
	if _, ok := (*m)[k1]; !ok {
		(*m)[k1] = make(map[K]T)
		(*m)[k1][k2] = x
	} else if existing, ok := (*m)[k1][k2]; !ok {
		(*m)[k1][k2] = x
	} else {
		(*m)[k1][k2] = min(x, existing)
	}

	return m
}

func (m *Dp[K, T]) SetMax(k1, k2 K, x T) *Dp[K, T] {
	if _, ok := (*m)[k1]; !ok {
		(*m)[k1] = make(map[K]T)
		(*m)[k1][k2] = x
	} else if existing, ok := (*m)[k1][k2]; !ok {
		(*m)[k1][k2] = x
	} else {
		(*m)[k1][k2] = max(x, existing)
	}

	return m
}
