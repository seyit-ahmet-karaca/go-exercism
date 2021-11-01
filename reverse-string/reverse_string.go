package reverse

func Reverse(input string) string {
	var reversed string
	for _,ch := range input {
		reversed = string(ch) + reversed
	}
	return reversed
}