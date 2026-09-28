func decodeString(s string) string {
	stringStack, numStack := []string{}, []int{}
	currentString, repeat := "", 0

	for _, c := range s {

		if c >= '0' && c <= '9' {

			// Build the repeat number.
			//
			// Example:
			// "123" → 1 → 12 → 123
			repeat = repeat*10 + int(c-'0')

		} else if c == '[' {

			// Save how many times the string inside
			// these brackets needs to be repeated.
			numStack = append(numStack, repeat)

			// Save the string that existed before '['.
			//
			// Example:
			// "ab2[c]"
			//
			// Save "ab" so we can restore it after
			// decoding "c".
			stringStack = append(stringStack, currentString)

			// Start building the string inside the brackets.
			currentString = ""

			// The repeat number has now been saved.
			repeat = 0

		} else if c == ']' {

			// Get the repeat count for this bracket pair.
			prevNum := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			// Get the string that existed before '['.
			prevString := stringStack[len(stringStack)-1]
			stringStack = stringStack[:len(stringStack)-1]

			// Save the string we just decoded inside
			// the brackets.
			//
			// Example:
			// 2[cd]
			// currentString = "cd"
			temp := currentString

			// Restore the string that existed before '['.
			//
			// Example:
			// "ab2[cd]"
			//
			// currentString becomes "ab".
			currentString = prevString

			// Append the decoded inner string repeatedly.
			//
			// Example:
			// prevString = "ab"
			// temp       = "cd"
			// prevNum    = 2
			//
			// result:
			// "ab" + "cd" + "cd"
			// = "abcdcd"
			for i := 0; i < prevNum; i++ {
				currentString += temp
			}

		} else {

			// Normal character.
			// Add it to the current string.
			currentString += string(c)
		}
	}

	return currentString
}