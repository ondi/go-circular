//
//
//

package circular

type List_t struct {
	root []interface{}
	head int
	tail int
	limit int
	written int
}

func New(limit int) (self * List_t) {
	self = &List_t{}
	self.root = make([]interface{}, limit)
	self.limit = limit
	return
}

func (self * List_t) PushBack(value interface{}) bool {
	if self.written == self.limit {
		return false
	}
	if self.tail == self.limit {
		self.tail = 0
	}
	self.root[self.tail] = value
	self.tail++
	self.written++
	return true
}

func (self * List_t) PopFront() (value interface{}, ok bool) {
	if self.written == 0 {
		return
	}
	if self.head == self.limit {
		self.head = 0
	}
	value = self.root[self.head]
	self.head++
	self.written--
	ok = true
	return
}

func (self * List_t) PushFront(value interface{}) bool {
	if self.written == self.limit {
		return false
	}
	if self.head == 0 {
		self.head = self.limit - 1
	} else {
		self.head--
	}
	self.root[self.head] = value
	self.written++
	return true
}

func (self * List_t) PopBack() (interface{}, bool) {
	if self.written == 0 {
		return nil, false
	}
	if self.tail == 0 {
		self.tail = self.limit - 1
	} else {
		self.tail--
	}
	self.written--
	return self.root[self.tail], true
}

func (self * List_t) RangeFront(f func(interface{}) bool) {
	cur := self.head
	for i := 0; i < self.written; i++ {
		if cur == self.limit {
			cur = 0
		}
		if f(self.root[cur]) == false {
			return
		}
		cur++
	}
}

func (self * List_t) RangeBack(f func(interface{}) bool) {
	cur := self.tail
	for i := 0; i < self.written; i++ {
		if cur == 0 {
			cur = self.limit - 1
		} else {
			cur--
		}
		if f(self.root[cur]) == false {
			return
		}
	}
}

func (self * List_t) Size() int {
	return self.written
}
