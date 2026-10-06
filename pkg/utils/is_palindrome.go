package utils

func IsPalindrome(s string) bool {
	runeSlice := []rune(s)
	length := len(runeSlice)
	for i := 0; i < length/2; i++ {
		if runeSlice[i] != runeSlice[length-1-i] {
			return false
		}
	}
	return true
}
