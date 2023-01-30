package multiset

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
	return &Multiset[T]{items: make(map[T]int), count: n}
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
