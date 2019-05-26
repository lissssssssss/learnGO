package main
import (
	"fmt"
)

func main() {
	letters := " "
	fmt.Println(isPalindrome(letters))
}
func isPalindrome(s string) bool {
	i, j := 0, len(s) - 1
	for i <= j {
		for i <= j && !(s[i] >= 'a' && s[i] <= 'z' ||s[i] >= 'A' && s[i] <= 'Z' || s[i] >= '0' && s[i] <= '9') {
			i++
		}
		for i <= j && !(s[j] >= 'a' && s[j] <= 'z' || s[j] >= 'A' && s[j] <= 'Z' || s[j] >= '0' && s[j] <= '9') {
			j--
		}
		if i > j {
			break
		}
		letterLeft := s[i]
		letterRight := s[j]
		if letterLeft >= 'A' && letterLeft <= 'Z' {
			letterLeft = letterLeft + 'a' - 'A'
		}
		if letterRight >= 'A' && letterRight <= 'Z' {
			letterRight = letterRight + 'a' - 'A'
		}

		if letterLeft != letterRight{
			fmt.Println(i,j)
			fmt.Println(s[i:i+1],s[j:j+1])
			return false
		}
		i++
		j--
	}
	return true
}
