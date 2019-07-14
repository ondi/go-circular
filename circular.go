//
//
//

package circular

type Circular_t struct {
	root []interface{}
	head int
	tail int
	limit int
	written int
}

func New(limit int) (self * Circular_t) {
	self = &Circular_t{}
	self.root = make([]interface{}, limit)
	self.limit = limit
	return
}

func (self * Circular_t) PushBack(value interface{}) bool {
	if self.written == self.limit {
		return false
	}
	if self.head == self.limit {
		self.head = 0
	}
	self.root[self.head] = value
	self.head++
	self.written++
	return true
}

func (self * Circular_t) PopBack() (interface{}, bool) {
	if self.written == 0 {
		return nil, false
	}
	if self.head == 0 {
		self.head = self.limit - 1
	} else {
		self.head--
	}
	self.written--
	return self.root[self.head], true
}

func (self * Circular_t) PushFront(value interface{}) bool {
	if self.written == self.limit {
		return false
	}
	if self.tail == 0 {
		self.tail = self.limit - 1
	} else {
		self.tail--
	}
	self.root[self.tail] = value
	self.written++
	return true
}

func (self * Circular_t) PopFront() (value interface{}, ok bool) {
	if self.written == 0 {
		return
	}
	if self.tail == self.limit {
		self.tail = 0
	}
	value = self.root[self.tail]
	self.tail++
	self.written--
	ok = true
	return
}

func (self * Circular_t) Size() int {
	return self.written
}
