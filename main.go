package main

import (
	"fmt"
	"strconv"
	_ "strconv"
)

type Stack []string

// IsEmpty: check if stack is empty
func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}

// Push a new value onto the stack
func (st *Stack) Push(str string) {
	*st = append(*st, str) //Simply append the new value to the end of the stack
}

// Pop Remove top element of stack. Return false if stack is empty.
func (st *Stack) Pop() bool {
	if st.IsEmpty() {
		return false
	} else {
		index := len(*st) - 1 // Get the index of top most element.
		*st = (*st)[:index]   // Remove it from the stack by slicing it off.
		return true
	}
}

// Top Return top element of stack. Return false if stack is empty.
func (st *Stack) Top() string {
	if st.IsEmpty() {
		return ""
	} else {
		index := len(*st) - 1   // Get the index of top most element.
		element := (*st)[index] // Index onto the slice and obtain the element.
		return element
	}
}

// Function to return precedence of operators
func prec(s string) int {
	if s == "^" {
		return 3
	} else if (s == "/") || (s == "*") {
		return 2
	} else if (s == "+") || (s == "-") {
		return 1
	} else {
		return -1
	}
}

func infixToPostfix(infix string) string {
	var sta Stack
	var postfix string
	for _, char := range infix {
		opchar := string(char)
		// if scanned character is operand, add it to output string
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			postfix = postfix + opchar
		} else if char == '(' {
			sta.Push(opchar)
		} else if char == ')' {
			for sta.Top() != "(" {
				postfix = postfix + sta.Top()
				sta.Pop()
			}
			sta.Pop()
		} else {
			for !sta.IsEmpty() && prec(opchar) <= prec(sta.Top()) {
				postfix = postfix + sta.Top()
				sta.Pop()
			}
			sta.Push(opchar)
		}
	}
	// Pop all the remaining elements from the stack
	for !sta.IsEmpty() {
		postfix = postfix + sta.Top()
		sta.Pop()
	}
	return postfix
}

// Stack Start From Here

type StackInt struct {
	s []int
}

// IsEmpty isEmpty() function
func (s *StackInt) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

// Length length() function
func (s *StackInt) Length() int {
	length := len(s.s)
	return length
}

// Print Print() function
func (s *StackInt) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

// Push Push() function
func (s *StackInt) Push(value int) {
	s.s = append(s.s, value)
}

// Pop Pop() function
func (s *StackInt) Pop() int {
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

// Top Top() function
func (s *StackInt) Top() int {
	length := len(s.s)
	res := s.s[length-1]
	return res
}

// stack End Here

func main() {
	var stack StackInt                 // create a stack variable of type Stack
	fmt.Print("Enter Your Equation: ") //Print function is used to display output in same line
	var infix string
	fmt.Scanln(&infix) // Taking input from user
	postfix := infixToPostfix(infix)
	//fmt.Printf("%s infix has %s postfix \n", infix, postfix)
	for i := 0; i < len(postfix); i++ {
		if postfix[i] != '+' && postfix[i] != '-' && postfix[i] != '/' && postfix[i] != '*' {
			int1, _ := strconv.ParseInt(string(postfix[i]), 6, 12)
			stack.Push(int(int1))
		} else {
			if stack.Length() >= 2 {
				a := stack.Pop() // Remove an item
				b := stack.Pop() // Remove an item
				if postfix[i] == '+' {
					stack.Push(a + b) // Add an item
				} else if postfix[i] == '-' {
					stack.Push(b - a) // Add an item
				} else if postfix[i] == '*' {
					stack.Push(a * b) // Add an item
				} else {
					stack.Push(b / a) // Add an item
				}
			}
		}
	}
	fmt.Println(stack.Pop())
}
