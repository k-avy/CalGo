package prefix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/k-avy/CalGo/pkg/stack"
)

func SolvePrefixExpression(str string) (int, error) {
	st := stack.Stack[int]()
	ele := strings.Fields(str)
	l := len(ele)
	tmp := make([]string, l)
	copy(tmp, ele)
	fmt.Println(tmp)
	for i := l - 1; i >= 0; i-- {
		num, err := strconv.Atoi(ele[i])
		tmp = tmp[:i]
		if err == nil {
			st.Push(num)
		} else {
			o1, err1 := st.Pop()
			o2, err1 := st.Pop()
			if err1 != nil {
				return -1, err1
			}
			if ele[i] == "+" {
				num = o1 + o2
			} else if ele[i] == "-" {
				num = o1 - o2
			} else if ele[i] == "*" {
				num = o1 * o2
			} else if ele[i] == "/" {
				num = o1 / o2
			} else {
				return -1, errors.New("Invalid Expression")
			}
			st.Push(num)
			fmt.Println(tmp, num)
		}
	}
	return st.Pop()
}

func ConvertPretoPost(expr string) (string, error) {
	st := stack.Stack[string]()
	elem := strings.Fields(expr)
	l := len(elem)
	tmp := make([]string, l)
	copy(tmp, elem)
	fmt.Println(tmp)
	for i := l - 1; i >= 0; i-- {
		tmp = tmp[:i]
		if elem[i] == "+" || elem[i] == "-" || elem[i] == "*" || elem[i] == "/" {
			op2, err := st.Pop()
			op1, err := st.Pop()
			if err != nil {
				return "nil", err
			}
			temp := string(op2) + " " + string(op1) + " " + string(elem[i])
			st.Push(temp)
		} else {
			st.Push(string(elem[i]))
			fmt.Println(tmp, elem[i])
		}
	}
	return st.Pop()
}

func ConvertPretoIn(s string) (string, error) {
	st := stack.Stack[string]()
	elem := strings.Fields(s)
	l := len(elem)
	t := make([]string, l)
	copy(t, elem)
	for i := l - 1; i >= 0; i-- {
		t = t[:i]
		if elem[i] == "+" || elem[i] == "-" || elem[i] == "*" || elem[i] == "/" {
			op2, err1 := st.Pop()
			op1, err1 := st.Pop()
			if err1 != nil {
				return "nil", err1
			}
			temp := string(op2 + " " + elem[i] + " " + op1)
			st.Push(temp)
		} else {
			st.Push(string(elem[i]))
		}
		st.Print()
	}
	return st.Pop()
}
