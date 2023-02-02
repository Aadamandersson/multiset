package multiset_test

import (
	"fmt"
	"testing"

	"github.com/aadamandersson/multiset"
)

func ExampleNew() {
	_ = multiset.New[int]()
}

func ExampleWithCapacity() {
	_ = multiset.WithCapacity[int](10)
}

func ExampleMultiset_Insert() {
	ms := multiset.New[int]()
	fmt.Println(ms.Insert(10))
	fmt.Println(ms.Insert(10))
	fmt.Println(ms.Insert(20))
	// Output:
	// 0
	// 1
	// 0
}

func ExampleMultiset_InsertMany() {
	ms := multiset.New[int]()
	fmt.Println(ms.InsertMany(10, 2))
	fmt.Println(ms.InsertMany(10, 4))
	fmt.Println(ms.InsertMany(20, 6))
	// Output:
	// 0
	// 2
	// 0
}

func ExampleMultiset_Union() {
	ms1 := multiset.New[string]()
	ms1.InsertMany("a", 2)
	ms1.InsertMany("b", 3)
	ms1.InsertMany("c", 1)
	ms1.InsertMany("d", 2)

	ms2 := multiset.New[string]()
	ms2.InsertMany("b", 2)
	ms2.InsertMany("c", 3)
	ms2.InsertMany("d", 2)
	ms2.InsertMany("e", 1)
	fmt.Println(ms1.Union(ms2))
	// Output:
	// Multiset{a:2, b:3, c:3, d:2, e:1}
}

func ExampleMultiset_Intersection() {
	ms1 := multiset.New[string]()
	ms1.InsertMany("a", 3)
	ms1.InsertMany("b", 2)
	ms1.InsertMany("c", 3)
	ms1.InsertMany("d", 1)

	ms2 := multiset.New[string]()
	ms2.InsertMany("a", 2)
	ms2.InsertMany("b", 1)
	ms2.InsertMany("c", 1)
	fmt.Println(ms1.Intersection(ms2))
	// Output:
	// Multiset{a:2, b:1, c:1}
}

func ExampleMultiset_Sum() {
	ms1 := multiset.New[string]()
	ms1.InsertMany("a", 2)
	ms1.InsertMany("b", 3)
	ms1.InsertMany("d", 1)

	ms2 := multiset.New[string]()
	ms2.InsertMany("a", 1)
	ms2.InsertMany("b", 3)
	ms2.InsertMany("c", 2)
	fmt.Println(ms1.Sum(ms2))
	// Output:
	// Multiset{a:3, b:6, c:2, d:1}
}

func ExampleMultiset_Difference() {
	ms1 := multiset.New[string]()
	ms1.InsertMany("a", 3)
	ms1.InsertMany("b", 2)
	ms1.InsertMany("c", 1)

	ms2 := multiset.New[string]()
	ms2.InsertMany("a", 2)
	ms2.InsertMany("b", 1)
	ms2.InsertMany("c", 1)
	ms2.InsertMany("d", 1)
	result := ms1.Difference(ms2)
	fmt.Println(result)
	fmt.Println(result.Len())
	// Output:
	// Multiset{a:1, b:1}
	// 2
}

func ExampleMultiset_Replace() {
	ms := multiset.New[int]()
	ms.InsertMany(10, 2)
	fmt.Println(ms.Get(10))
	fmt.Println(ms.Replace(10))
	fmt.Println(ms.Get(10))
	// Output:
	// 10 2 true
	// 2
	// 10 1 true
}

func ExampleMultiset_Remove() {
	ms := multiset.New[int]()
	ms.InsertMany(10, 2)
	fmt.Println(ms.Contains(10))
	ms.Remove(10)
	fmt.Println(ms.Contains(10))
	ms.Remove(10)
	fmt.Println(ms.Contains(10))
	// Output:
	// 2
	// 1
	// 0
}

func ExampleMultiset_Get() {
	ms := multiset.New[int]()
	ms.Insert(10)
	ms.Insert(20)
	ms.Insert(20)
	fmt.Println(ms.Get(10))
	fmt.Println(ms.Get(20))
	fmt.Println(ms.Get(30))
	// Output:
	// 10 1 true
	// 20 2 true
	// 0 0 false
}

func ExampleMultiset_Contains() {
	ms := multiset.New[int]()
	ms.Insert(10)
	ms.Insert(10)
	ms.Insert(20)
	fmt.Println(ms.Contains(10))
	fmt.Println(ms.Contains(20))
	fmt.Println(ms.Contains(40))
	// Output:
	// 2
	// 1
	// 0
}

func ExampleMultiset_IsEmpty() {
	ms := multiset.New[int]()
	fmt.Println(ms.IsEmpty())
	ms.Insert(10)
	fmt.Println(ms.IsEmpty())
	// Output:
	// true
	// false
}

func ExampleMultiset_Len() {
	ms := multiset.New[int]()
	ms.Insert(10)
	ms.Insert(10)
	ms.Insert(20)
	fmt.Println(ms.Len())
	// Output:
	// 3
}

func ExampleMultiset_Cardinality() {
	ms := multiset.New[int]()
	ms.InsertMany(10, 1)
	ms.InsertMany(20, 2)
	ms.InsertMany(30, 3)
	fmt.Println(ms.Cardinality())
	// Output:
	// 3
}

func ExampleMultiset_Each() {
	ms1 := multiset.New[int]()
	ms1.Insert(10)
	ms1.Insert(20)
	ms1.InsertMany(30, 2)

	ms2 := multiset.New[int]()
	ms1.Each(func(value, multiplicity int) bool {
		ms2.InsertMany(value, multiplicity)
		return false
	})
	fmt.Println(ms1.Equal(ms2))
	// Output:
	// true
}

func ExampleMultiset_Clone() {
	ms1 := multiset.New[int]()
	ms1.InsertMany(20, 2)
	ms1.InsertMany(30, 3)

	ms2 := ms1.Clone()
	fmt.Println(ms1.Equal(ms2))
	ms2.InsertMany(40, 4)
	fmt.Println(ms1.Equal(ms2))
	// Output:
	// true
	// false
}

func TestEqual(t *testing.T) {
	ms1 := multiset.New[int]()
	ms2 := multiset.New[int]()

	ms1.InsertMany(10, 2)
	ms2.InsertMany(10, 2)

	ms1.Insert(20)
	ms2.Insert(20)

	if !ms1.Equal(ms2) {
		t.Errorf("Multisets are not equal: left %v, right %v\n", ms1, ms2)
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		fn   func() string
		want string
	}{
		{
			fn: func() string {
				ms := multiset.New[int]()
				return ms.String()
			},
			want: "Multiset{}",
		},
		{
			fn: func() string {
				ms := multiset.New[int]()
				ms.InsertMany(1, 2)
				ms.InsertMany(2, 3)
				return ms.String()
			},
			want: "Multiset{1:2, 2:3}",
		},
		{
			fn: func() string {
				ms := multiset.New[int]()
				ms.InsertMany(2, 3)
				ms.InsertMany(1, 2)
				ms.InsertMany(3, 4)
				return ms.String()
			},
			want: "Multiset{1:2, 2:3, 3:4}",
		},
	}

	for _, c := range cases {
		got := c.fn()
		if got != c.want {
			t.Errorf("String() = %s, want %s", got, c.want)
		}
	}
}

func TestUnion(t *testing.T) {
	cases := []struct {
		gotFn  func() *multiset.Multiset[string]
		wantFn func() *multiset.Multiset[string]
	}{
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 3)
				ms1.InsertMany("b", 2)

				ms2 := multiset.New[string]()
				ms2.InsertMany("a", 2)
				ms2.InsertMany("b", 3)
				return ms1.Union(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				ms := multiset.New[string]()
				ms.InsertMany("a", 3)
				ms.InsertMany("b", 3)
				return ms
			},
		},
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 3)
				ms1.InsertMany("b", 4)

				ms2 := multiset.New[string]()
				return ms1.Union(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				ms := multiset.New[string]()
				ms.InsertMany("a", 3)
				ms.InsertMany("b", 4)
				return ms
			},
		},
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 3)

				var ms2 *multiset.Multiset[string]
				return ms1.Union(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				ms := multiset.New[string]()
				ms.InsertMany("a", 3)
				return ms
			},
		},
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 2)
				ms1.InsertMany("b", 3)
				ms1.InsertMany("c", 1)
				ms1.InsertMany("d", 2)

				ms2 := multiset.New[string]()
				ms2.InsertMany("b", 2)
				ms2.InsertMany("c", 3)
				ms2.InsertMany("d", 2)
				ms2.InsertMany("e", 1)
				return ms1.Union(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				ms := multiset.New[string]()
				ms.InsertMany("a", 2)
				ms.InsertMany("b", 3)
				ms.InsertMany("c", 3)
				ms.InsertMany("d", 2)
				ms.InsertMany("e", 1)
				return ms
			},
		},
	}

	for _, c := range cases {
		got := c.gotFn()
		want := c.wantFn()
		if !got.Equal(want) {
			t.Errorf("Len() = %d, want %d\n", got.Len(), want.Len())
			t.Errorf("Union() = %v, want %v", got, want)
		}
	}
}

func TestIntersection(t *testing.T) {
	cases := []struct {
		gotFn  func() *multiset.Multiset[string]
		wantFn func() *multiset.Multiset[string]
	}{
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 3)
				ms1.InsertMany("b", 2)

				ms2 := multiset.New[string]()
				ms2.InsertMany("a", 2)
				ms2.InsertMany("b", 3)
				return ms1.Intersection(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				ms := multiset.New[string]()
				ms.InsertMany("a", 2)
				ms.InsertMany("b", 2)
				return ms
			},
		},
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 3)

				ms2 := multiset.New[string]()
				return ms1.Intersection(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				return multiset.New[string]()
			},
		},
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 3)

				var ms2 *multiset.Multiset[string]
				return ms1.Intersection(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				return multiset.New[string]()
			},
		},
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 2)
				ms1.InsertMany("b", 3)
				ms1.InsertMany("c", 1)
				ms1.InsertMany("d", 2)

				ms2 := multiset.New[string]()
				ms2.InsertMany("b", 2)
				ms2.InsertMany("c", 3)
				ms2.InsertMany("d", 2)
				ms2.InsertMany("e", 1)
				return ms1.Intersection(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				ms := multiset.New[string]()
				ms.InsertMany("b", 2)
				ms.InsertMany("c", 1)
				ms.InsertMany("d", 2)
				return ms
			},
		},
	}

	for _, c := range cases {
		got := c.gotFn()
		want := c.wantFn()
		if !got.Equal(want) {
			t.Errorf("Len() = %d, want %d\n", got.Len(), want.Len())
			t.Errorf("Intersection() = %v, want %v", got, want)
		}
	}
}

func TestSum(t *testing.T) {
	cases := []struct {
		gotFn  func() *multiset.Multiset[string]
		wantFn func() *multiset.Multiset[string]
	}{
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 3)
				ms1.InsertMany("b", 2)

				ms2 := multiset.New[string]()
				ms2.InsertMany("a", 2)
				ms2.InsertMany("b", 3)
				return ms1.Sum(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				ms := multiset.New[string]()
				ms.InsertMany("a", 5)
				ms.InsertMany("b", 5)
				return ms
			},
		},
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 3)
				ms1.InsertMany("b", 4)

				ms2 := multiset.New[string]()
				return ms1.Sum(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				ms := multiset.New[string]()
				ms.InsertMany("a", 3)
				ms.InsertMany("b", 4)
				return ms
			},
		},
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 3)

				var ms2 *multiset.Multiset[string]
				return ms1.Sum(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				ms := multiset.New[string]()
				ms.InsertMany("a", 3)
				return ms
			},
		},
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 2)
				ms1.InsertMany("b", 3)
				ms1.InsertMany("c", 1)
				ms1.InsertMany("d", 2)

				ms2 := multiset.New[string]()
				ms2.InsertMany("b", 2)
				ms2.InsertMany("c", 3)
				ms2.InsertMany("d", 2)
				ms2.InsertMany("e", 1)
				return ms1.Sum(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				ms := multiset.New[string]()
				ms.InsertMany("a", 2)
				ms.InsertMany("b", 5)
				ms.InsertMany("c", 4)
				ms.InsertMany("d", 4)
				ms.InsertMany("e", 1)
				return ms
			},
		},
	}

	for _, c := range cases {
		got := c.gotFn()
		want := c.wantFn()
		if !got.Equal(want) {
			t.Errorf("Len() = %d, want %d\n", got.Len(), want.Len())
			t.Errorf("Sum() = %v, want %v", got, want)
		}
	}
}

func TestDifference(t *testing.T) {
	cases := []struct {
		gotFn  func() *multiset.Multiset[string]
		wantFn func() *multiset.Multiset[string]
	}{
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 3)
				ms1.InsertMany("b", 3)

				ms2 := multiset.New[string]()
				ms2.InsertMany("a", 2)
				ms2.InsertMany("b", 2)
				return ms1.Difference(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				ms := multiset.New[string]()
				ms.InsertMany("a", 1)
				ms.InsertMany("b", 1)
				return ms
			},
		},
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 3)
				ms1.InsertMany("b", 4)

				ms2 := multiset.New[string]()
				return ms1.Difference(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				ms := multiset.New[string]()
				ms.InsertMany("a", 3)
				ms.InsertMany("b", 4)
				return ms
			},
		},
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 3)

				var ms2 *multiset.Multiset[string]
				return ms1.Difference(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				ms := multiset.New[string]()
				ms.InsertMany("a", 3)
				return ms
			},
		},
		{
			gotFn: func() *multiset.Multiset[string] {
				ms1 := multiset.New[string]()
				ms1.InsertMany("a", 2)
				ms1.InsertMany("b", 3)
				ms1.InsertMany("c", 1)
				ms1.InsertMany("d", 2)

				ms2 := multiset.New[string]()
				ms2.InsertMany("b", 2)
				ms2.InsertMany("c", 3)
				ms2.InsertMany("d", 2)
				ms2.InsertMany("e", 1)
				return ms1.Difference(ms2)
			},
			wantFn: func() *multiset.Multiset[string] {
				ms := multiset.New[string]()
				ms.InsertMany("a", 2)
				ms.InsertMany("b", 1)
				return ms
			},
		},
	}

	for _, c := range cases {
		got := c.gotFn()
		want := c.wantFn()
		if !got.Equal(want) {
			t.Errorf("Len() = %d, want %d\n", got.Len(), want.Len())
			t.Errorf("Difference() = %v, want %v", got, want)
		}
	}
}
