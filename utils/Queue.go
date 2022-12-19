package utils

type Queue []string

func (s *Queue) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Queue) Push(strSlice ...string) {
	for _, item := range strSlice {
		*s = append(*s, item) // Simply append the new value to the end of the stack
	}
}

func (s *Queue) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		element := (*s)[0]
		*s = (*s)[1:]
		return element, true
	}
}
func (s *Queue) Reverse() {
	reversed := Queue{}
	count := len(*s)
	for i := 0; i < count; i++ {
		crate, _ := s.Pop()
		reversed.Push(crate)
	}

	*s = reversed
}
