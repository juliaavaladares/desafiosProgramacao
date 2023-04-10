package easy

func IsValid(s string) bool {

	parentheses := map[string]string{")": "(", "]": "[", "}": "{"}
	stack := []string{}

	for _, c := range s {
		character := string(c)
		// if character is in list of close parentheses
		openParentheses, ok := parentheses[character]
		if ok {
			if (len(stack) > 0) && (stack[len(stack)-1] == openParentheses) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, character)
		}
	}

	return len(stack) == 0

}
