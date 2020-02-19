package reverse_string

// considering the case when the character's size in more than one byte like japaniese
// here we are dealing with runes instead of bytes
func ReverseString(input string) string {
	var reversed string

	for _, char := range input {
		reversed = string(char) + reversed
	}

	return reversed
}

//func ReverseString(input string) string {
//	var reversed strings.Builder
//
//	for i := len(input)-1; i >= 0; i-- {
//		fmt.Println(string(input[i]))
//		reversed.WriteByte(input[i])
//	}
//
//	return reversed.String()
//}