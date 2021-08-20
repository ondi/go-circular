//
//
//

package circular

type List_t struct {
	root    []interface{}
	front   int
	back    int
	written int
}

func New(limit int) (self *List_t) {
	self = &List_t{}
	self.root = make([]interface{}, limit)
	return
}

func (self *List_t) Back() (interface{}, bool) {
	if self.written == 0 {
		return nil, false
	}
	if self.back == 0 {
		return self.root[len(self.root)-1], true
	}
	return self.root[self.back-1], true
}

func (self *List_t) Front() (interface{}, bool) {
	if self.written == 0 {
		return nil, false
	}
	if self.front == len(self.root) {
		return self.root[0], true
	}
	return self.root[self.front], true
}

func (self *List_t) PushBack(value interface{}) bool {
	if self.written == len(self.root) {
		return false
	}
	if self.back == len(self.root) {
		self.back = 0
	}
	self.root[self.back] = value
	self.back++
	self.written++
	return true
}

func (self *List_t) PopFront() (value interface{}, ok bool) {
	if self.written == 0 {
		return
	}
	if self.front == len(self.root) {
		self.front = 0
	}
	value = self.root[self.front]
	self.front++
	self.written--
	ok = true
	return
}

func (self *List_t) PushFront(value interface{}) bool {
	if self.written == len(self.root) {
		return false
	}
	if self.front == 0 {
		self.front = len(self.root) - 1
	} else {
		self.front--
	}
	self.root[self.front] = value
	self.written++
	return true
}

func (self *List_t) PopBack() (interface{}, bool) {
	if self.written == 0 {
		return nil, false
	}
	if self.back == 0 {
		self.back = len(self.root) - 1
	} else {
		self.back--
	}
	self.written--
	return self.root[self.back], true
}

func (self *List_t) RangeFront(f func(interface{}) bool) {
	cur := self.front
	for i := 0; i < self.written; i++ {
		if cur == len(self.root) {
			cur = 0
		}
		if f(self.root[cur]) == false {
			return
		}
		cur++
	}
}

func (self *List_t) RangeBack(f func(interface{}) bool) {
	cur := self.back
	for i := 0; i < self.written; i++ {
		if cur == 0 {
			cur = len(self.root) - 1
		} else {
			cur--
		}
		if f(self.root[cur]) == false {
			return
		}
	}
}

func (self *List_t) Size() int {
	return self.written
}
