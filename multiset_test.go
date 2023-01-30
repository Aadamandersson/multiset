package multiset_test

import (
	"fmt"

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
