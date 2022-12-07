package utils

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
func (s *Stack) Reverse() {
	reversed := Stack{}
	count := len(*s)
	for i := 0; i < count; i++ {
		crate, _ := s.Pop()
		reversed.Push(crate)
	}

	*s = reversed
}
