package main

import "fmt"

// Основная структура стека
type Stack struct {
	elemes []rune
}

// Метод для добавления в стек нового элемента
func (s *Stack) push(a rune) {
	s.elemes = append(s.elemes, a)
}

// Метод для получения последнего элемента и последующего удаления из стека
func (s *Stack) pop() rune {
	element := s.elemes[len(s.elemes)-1]
	s.elemes = s.elemes[:len(s.elemes)-1]
	return element
}

// Метод получения последнего элемента в стеке
func (s *Stack) top() rune {
	return s.elemes[len(s.elemes)-1]
}

func reverseParentheses(s string) string {
	st := Stack{}
	for _, char := range s {
		if char == ')' {
			for st.top() != '(' {
				fmt.Print(string(st.pop()))
			}
			st.pop()
		} else {
			st.push(char)
		}
	}
	return ""
}

func main() {
	reverseParentheses("(u(love)i)")
	// fmt.Println(s)
}
