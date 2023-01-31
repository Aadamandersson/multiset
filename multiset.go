package multiset

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
)

type Multiset[T comparable] struct {
	items map[T]int
	count int
}

// New contructs an empty multiset.
//
// The multiset is created with an initial capacity of 0.
// The initial capacity does not bound its size,
// and it grows to accomodate the number of values stored in it.
func New[T comparable]() *Multiset[T] {
	return &Multiset[T]{items: make(map[T]int), count: 0}
}

// WithCapacity constructs an empty multiset with the specified capacity.
//
// The multiset is created with an initial capacity of n.
// The initial capacity does not bound its size,
// and it grows to accomodate the number of values stored in it.
func WithCapacity[T comparable](n int) *Multiset[T] {
	return &Multiset[T]{items: make(map[T]int, n), count: 0}
}

// Insert inserts a new value v to multiset m.
//
// Insert returns the number of occurences of value v previously in
// multiset m.
func (m *Multiset[T]) Insert(v T) int {
	return m.InsertMany(v, 1)
}

// InsertMany inserts value v to multiset m, with n number of occurences.
//
// InsertMany returns the number of occurences of value v previously
// in multiset m.
func (m *Multiset[T]) InsertMany(v T, n int) int {
	if n == 0 {
		return m.Contains(v)
	}

	m.count += n
	pn, ok := m.items[v]
	if ok {
		m.items[v] = pn + n
	} else {
		m.items[v] = n
	}
	return pn
}

// Union constructs a new multiset union of multiset m and other.
//
// The resulting multiset is a multiset of the maximum multiplicity
// of items present in m and other.
func (m *Multiset[T]) Union(other *Multiset[T]) *Multiset[T] {
	result := WithCapacity[T](max(len(m.items), len(other.items)))
	m.Each(func(v T, n int) bool {
		result.InsertMany(v, n)
		return false
	})

	other.Each(func(otherV T, otherN int) bool {
		if n, ok := result.items[otherV]; !ok || otherN > n {
			result.count -= n
			result.items[otherV] = otherN
			result.count += otherN
		}
		return false
	})

	return result
}

// Replace replaces all existing occurences of value v in multiset m, if any, with 1.
//
// Replace returns the number of occurences of value v previously in
// multiset m.
func (m *Multiset[T]) Replace(v T) int {
	n := m.items[v]
	m.count -= n
	m.items[v] = 1
	m.count += 1
	return n
}

// Remove removes value v from multiset m.
//
// Remove returns the number of occurences of value v previously in
// multiset m.
func (m *Multiset[T]) Remove(v T) int {
	n, ok := m.items[v]
	if !ok {
		return 0
	}

	m.count -= 1
	if n == 1 {
		delete(m.items, v)
	} else {
		m.items[v] = n - 1
	}

	return n
}

// Get returns a value from multiset m that equals value v,
// along with its number of occurences and a boolean that is true.
//
// Get returns the zero value of type T, count 0 and false if value v
// does not exist in multiset m.
func (m *Multiset[T]) Get(v T) (T, int, bool) {
	for val, c := range m.items {
		if v == val {
			return val, c, true
		}
	}
	return *new(T), 0, false
}

// Contains returns the number of occurences of value v in multiset m.
func (m *Multiset[T]) Contains(v T) int {
	if n, ok := m.items[v]; ok {
		return n
	}
	return 0
}

// IsEmpty returns true if there are no items in multiset m,
// otherwise false.
func (m *Multiset[T]) IsEmpty() bool {
	return m.count == 0
}

// Len returns the number of items of multiset m.
//
// Duplicates are counted.
func (m *Multiset[T]) Len() int {
	return m.count
}

// Each iterates over all items and calls f for each item present in
// multiset m.
//
// If f returns true, Each stops the iteration.
func (m *Multiset[T]) Each(f func(T, int) bool) {
	for v, n := range m.items {
		if f(v, n) {
			break
		}
	}
}

// Equal returns true if the length and items of
// multiset m and other are equal.
//
// TODO: Remove package "golang.org/x/exp/maps" when
// https://github.com/golang/go/issues/57436 is accepted.
func (m *Multiset[T]) Equal(other *Multiset[T]) bool {
	return m.Len() == other.Len() && maps.Equal(m.items, other.items)
}

// String returns a formatted multiset with the following format:
//
// Multiset{1: 2, 2: 3}
func (m *Multiset[T]) String() string {
	items := make([]string, 0, len(m.items))

	m.Each(func(v T, n int) bool {
		items = append(items, fmt.Sprintf("%v:%d", v, n))
		return false
	})

	sort.Strings(items)
	return fmt.Sprintf("Multiset{%s}", strings.Join(items, ", "))
}

// max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
