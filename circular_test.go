package circular

import "fmt"
import "testing"

func ExampleCircular1() {
	var value interface{}
	var ok bool
	
	c := New(2)
	
	ok = c.PushBack(1)
	fmt.Printf("%v\n", ok)
	
	ok = c.PushBack(2)
	fmt.Printf("%v\n", ok)
	
	ok = c.PushBack(3)
	fmt.Printf("%v\n", ok)
	
	value, ok = c.PopBack()
	fmt.Printf("%v %v\n", value, ok)
	
	value, ok = c.PopBack()
	fmt.Printf("%v %v\n", value, ok)
	
	value, ok = c.PopBack()
	fmt.Printf("%v %v\n", value, ok)
/* Output:
true
true
false
2 true
1 true
<nil> false
*/
}

func ExampleCircular2() {
	var value interface{}
	var ok bool
	
	c := New(2)
	
	ok = c.PushFront(1)
	fmt.Printf("%v\n", ok)
	
	ok = c.PushFront(2)
	fmt.Printf("%v\n", ok)
	
	ok = c.PushFront(3)
	fmt.Printf("%v\n", ok)
	
	value, ok = c.PopFront()
	fmt.Printf("%v %v\n", value, ok)
	
	value, ok = c.PopFront()
	fmt.Printf("%v %v\n", value, ok)
	
	value, ok = c.PopFront()
	fmt.Printf("%v %v\n", value, ok)
/* Output:
true
true
false
2 true
1 true
<nil> false
*/
}

func ExampleCircular3() {
	var value interface{}
	var ok bool
	
	c := New(2)
	
	ok = c.PushBack(1)
	fmt.Printf("%v\n", ok)
	
	ok = c.PushBack(2)
	fmt.Printf("%v\n", ok)
	
	ok = c.PushBack(3)
	fmt.Printf("%v\n", ok)
	
	value, ok = c.PopFront()
	fmt.Printf("%v %v\n", value, ok)
	
	value, ok = c.PopFront()
	fmt.Printf("%v %v\n", value, ok)
	
	value, ok = c.PopFront()
	fmt.Printf("%v %v\n", value, ok)
/* Output:
true
true
false
1 true
2 true
<nil> false
*/
}

func ExampleCircular4() {
	var value interface{}
	var ok bool
	
	c := New(2)
	
	ok = c.PushFront(1)
	fmt.Printf("%v\n", ok)
	
	ok = c.PushFront(2)
	fmt.Printf("%v\n", ok)
	
	ok = c.PushFront(3)
	fmt.Printf("%v\n", ok)
	
	value, ok = c.PopBack()
	fmt.Printf("%v %v\n", value, ok)
	
	value, ok = c.PopBack()
	fmt.Printf("%v %v\n", value, ok)
	
	value, ok = c.PopBack()
	fmt.Printf("%v %v\n", value, ok)
/* Output:
true
true
false
1 true
2 true
<nil> false
*/
}

func ExampleCircular5() {
	var value interface{}
	var ok bool
	
	c := New(2)
	
	ok = c.PushBack(1)
	fmt.Printf("%v\n", ok)
	
	ok = c.PushFront(2)
	fmt.Printf("%v\n", ok)
	
	ok = c.PushBack(3)
	fmt.Printf("%v\n", ok)
	
	value, ok = c.PopBack()
	fmt.Printf("%v %v\n", value, ok)
	
	value, ok = c.PopFront()
	fmt.Printf("%v %v\n", value, ok)
	
	value, ok = c.PopBack()
	fmt.Printf("%v %v\n", value, ok)
/* Output:
true
true
false
1 true
2 true
<nil> false
*/
}

func ExampleCircular6() {
	var value interface{}
	var ok bool
	
	c := New(2)
	
	ok = c.PushFront(1)
	fmt.Printf("%v\n", ok)
	
	ok = c.PushBack(2)
	fmt.Printf("%v\n", ok)
	
	ok = c.PushFront(3)
	fmt.Printf("%v\n", ok)
	
	value, ok = c.PopFront()
	fmt.Printf("%v %v\n", value, ok)
	
	value, ok = c.PopBack()
	fmt.Printf("%v %v\n", value, ok)
	
	value, ok = c.PopFront()
	fmt.Printf("%v %v\n", value, ok)
/* Output:
true
true
false
1 true
2 true
<nil> false
*/
}

func ExampleCircular7() {
	var value interface{}
	var ok bool
	
	c := New(2)
	
	ok = c.PushFront(1)
	fmt.Printf("%v\n", ok)
	
	value, ok = c.PopBack()
	fmt.Printf("%v %v\n", value, ok)
/* Output:
true
1 true
*/
}

func ExampleCircular8() {
	var value interface{}
	var ok bool
	
	c := New(2)
	
	ok = c.PushBack(1)
	fmt.Printf("%v\n", ok)
	
	value, ok = c.PopFront()
	fmt.Printf("%v %v\n", value, ok)
/* Output:
true
1 true
*/
}

func ExampleCircular9() {
	c := New(3)
	
	c.PushBack(1)
	c.PushBack(2)
	c.PushBack(3)
	c.RangeFront(func(v interface{}) bool {
		fmt.Printf("%v\n", v)
		return true
	})
	c.RangeBack(func(v interface{}) bool {
		fmt.Printf("%v\n", v)
		return true
	})
/* Output:
1
2
3
3
2
1
*/
}

func ExampleCircular10() {
	c := New(3)
	
	c.PushFront(3)
	c.PushFront(2)
	c.PushFront(1)
	c.RangeFront(func(v interface{}) bool {
		fmt.Printf("%v\n", v)
		return true
	})
	c.RangeBack(func(v interface{}) bool {
		fmt.Printf("%v\n", v)
		return true
	})
/* Output:
1
2
3
3
2
1
*/
}

func ExampleCircular11() {
	c := New(3)
	
	c.PushBack(4)
	c.PushFront(2)
	c.PopBack()
	c.PushFront(1)
	c.PushBack(3)
	c.RangeFront(func(v interface{}) bool {
		fmt.Printf("%v\n", v)
		return true
	})
	c.RangeBack(func(v interface{}) bool {
		fmt.Printf("%v\n", v)
		return true
	})
/* Output:
1
2
3
3
2
1
*/
}

func BenchmarkCircular1(b * testing.B) {
	b.ReportAllocs()
	c := New(b.N)
	for i := 0; i < b.N; i++{
		c.PushBack("lalala")
	}
}
