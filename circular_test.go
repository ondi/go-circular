package circular

import (
	"testing"

	"gotest.tools/assert"
)

func TestCircular1(t *testing.T) {
	var value int
	var ok bool

	c := New[int](2)

	ok = c.PushBack(1)
	assert.Assert(t, ok == true)

	ok = c.PushBack(2)
	assert.Assert(t, ok == true)

	ok = c.PushBack(3)
	assert.Assert(t, ok == false)

	value, ok = c.PopBack()
	assert.Assert(t, value == 2 && ok == true)

	value, ok = c.PopBack()
	assert.Assert(t, value == 1 && ok == true)

	value, ok = c.PopBack()
	assert.Assert(t, value == 0 && ok == false)
}

func TestCircular2(t *testing.T) {
	var value int
	var ok bool

	c := New[int](2)

	ok = c.PushFront(1)
	assert.Assert(t, ok == true)

	ok = c.PushFront(2)
	assert.Assert(t, ok == true)

	ok = c.PushFront(3)
	assert.Assert(t, ok == false)

	value, ok = c.PopFront()
	assert.Assert(t, value == 2 && ok == true)

	value, ok = c.PopFront()
	assert.Assert(t, value == 1 && ok == true)

	value, ok = c.PopFront()
	assert.Assert(t, value == 0 && ok == false)
}

func TestCircular3(t *testing.T) {
	var value int
	var ok bool

	c := New[int](2)

	ok = c.PushBack(1)
	assert.Assert(t, ok == true)

	ok = c.PushBack(2)
	assert.Assert(t, ok == true)

	ok = c.PushBack(3)
	assert.Assert(t, ok == false)

	value, ok = c.PopFront()
	assert.Assert(t, value == 1 && ok == true)

	value, ok = c.PopFront()
	assert.Assert(t, value == 2 && ok == true)

	value, ok = c.PopFront()
	assert.Assert(t, value == 0 && ok == false)
}

func TestCircular4(t *testing.T) {
	var value int
	var ok bool

	c := New[int](2)

	ok = c.PushFront(1)
	assert.Assert(t, ok == true)

	ok = c.PushFront(2)
	assert.Assert(t, ok == true)

	ok = c.PushFront(3)
	assert.Assert(t, ok == false)

	value, ok = c.PopBack()
	assert.Assert(t, value == 1 && ok == true)

	value, ok = c.PopBack()
	assert.Assert(t, value == 2 && ok == true)

	value, ok = c.PopBack()
	assert.Assert(t, value == 0 && ok == false)
}

func TestCircular5(t *testing.T) {
	var value int
	var ok bool

	c := New[int](2)

	ok = c.PushBack(1)
	assert.Assert(t, ok == true)

	ok = c.PushFront(2)
	assert.Assert(t, ok == true)

	ok = c.PushBack(3)
	assert.Assert(t, ok == false)

	value, ok = c.PopBack()
	assert.Assert(t, value == 1 && ok == true)

	value, ok = c.PopFront()
	assert.Assert(t, value == 2 && ok == true)

	value, ok = c.PopBack()
	assert.Assert(t, value == 0 && ok == false)
}

func TestCircular6(t *testing.T) {
	var value int
	var ok bool

	c := New[int](2)

	ok = c.PushFront(1)
	assert.Assert(t, ok == true)

	ok = c.PushBack(2)
	assert.Assert(t, ok == true)

	ok = c.PushFront(3)
	assert.Assert(t, ok == false)

	value, ok = c.PopFront()
	assert.Assert(t, value == 1 && ok == true)

	value, ok = c.PopBack()
	assert.Assert(t, value == 2 && ok == true)

	value, ok = c.PopFront()
	assert.Assert(t, value == 0 && ok == false)
}

func TestCircular7(t *testing.T) {
	var value int
	var ok bool

	c := New[int](2)

	ok = c.PushFront(1)
	assert.Assert(t, ok == true)

	value, ok = c.PopBack()
	assert.Assert(t, value == 1 && ok == true)
}

func TestCircular8(t *testing.T) {
	var value int
	var ok bool

	c := New[int](2)

	ok = c.PushBack(1)
	assert.Assert(t, ok == true)

	value, ok = c.PopFront()
	assert.Assert(t, value == 1 && ok == true)
}

func TestCircular9(t *testing.T) {
	c := New[int](3)

	c.PushBack(1)
	c.PushBack(2)
	c.PushBack(3)
	i := 1
	c.RangeFront(
		func(v int) bool {
			assert.Assert(t, v == i)
			i++
			return true
		},
	)
	i = 3
	c.RangeBack(
		func(v int) bool {
			assert.Assert(t, v == i)
			i--
			return true
		},
	)
}

func TestCircular10(t *testing.T) {
	c := New[int](3)

	c.PushFront(3)
	c.PushFront(2)
	c.PushFront(1)
	i := 1
	c.RangeFront(
		func(v int) bool {
			assert.Assert(t, v == i)
			i++
			return true
		},
	)
	i = 3
	c.RangeBack(
		func(v int) bool {
			assert.Assert(t, v == i)
			i--
			return true
		},
	)
}

func TestCircular11(t *testing.T) {
	c := New[int](3)

	c.PushBack(4)
	c.PushFront(2)
	c.PopBack()
	c.PushFront(1)
	c.PushBack(3)
	i := 1
	c.RangeFront(
		func(v int) bool {
			assert.Assert(t, v == i)
			i++
			return true
		},
	)
	i = 3
	c.RangeBack(
		func(v int) bool {
			assert.Assert(t, v == i)
			i--
			return true
		},
	)
}

func TestCircular12(t *testing.T) {
	var value int
	var ok bool

	c := New[int](2)

	ok = c.PushBack(1)
	assert.Assert(t, ok == true)

	value, ok = c.Back()
	assert.Assert(t, value == 1 && ok == true)

	value, ok = c.Front()
	assert.Assert(t, value == 1 && ok == true)

	ok = c.PushBack(2)
	assert.Assert(t, ok == true)

	value, ok = c.Back()
	assert.Assert(t, value == 2 && ok == true)

	value, ok = c.Front()
	assert.Assert(t, value == 1 && ok == true)

	ok = c.PushBack(3)
	assert.Assert(t, ok == false)

	value, ok = c.Back()
	assert.Assert(t, value == 2 && ok == true)

	value, ok = c.Front()
	assert.Assert(t, value == 1 && ok == true)
}

func BenchmarkCircular1(b *testing.B) {
	b.ReportAllocs()
	c := New[string](b.N)
	for i := 0; i < b.N; i++ {
		c.PushBack("lalala")
	}
}
