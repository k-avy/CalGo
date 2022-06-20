package stack

import "errors"

type stack[T any] struct {
	Push   func(T)
	Pop    func() (T, error)
	Length func() int
	Print  func()
}

func Stack[T any]() stack[T] {
	slice := make([]T, 0)
	return stack[T]{
		Push: func(t T) {
			slice = append(slice, t)
		},
		Pop: func() (T, error) {
			var ele T
			n := len(slice)
			if n == 0 {
				return ele, errors.New("Stack is empty")
			} else {
				ele = slice[n-1]
				slice = slice[:n-1]
				return ele, nil
			}
		},
		Length: func() int {
			return len(slice)
		},
		Print: func() {
			for _, val := range slice {
				print(val, " ")
			}
			println("")
		},
	}
}
