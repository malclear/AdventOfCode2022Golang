package utils

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) PeekAll() interface{} {
	return s.items
}

func (s *Stack) Contains(item interface{}) bool {
	for _, thisItem := range s.items {
		if item == thisItem {
			return true
		}
	}
	return false
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Reverse() {
	reversed := Stack{}
	count := len(s.items)
	for i := 0; i < count; i++ {
		t := s.Pop()
		reversed.Push(t)
	}

	*s = reversed
}
