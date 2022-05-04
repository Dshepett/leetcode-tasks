package easy

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// Input: s = "()[]{}"
// Output: true

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func isValid(s string) bool {
	var stack Stack
	for _, val := range s {
		char := string(val)
		if val == '[' || val == '{' || val == '(' {
			stack.Push(char)
		} else {
			par, empty := stack.Pop()
			if !empty {
				return false
			}
			if par == "[" && char != "]" || par == "{" && char != "}" || par == "(" && char != ")" {
				return false

			}
		}
	}
	return stack.IsEmpty()
}
