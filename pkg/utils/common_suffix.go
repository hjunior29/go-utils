package utils

func CommonSuffix(a, b string) string {
	runeseA := []rune(a)
	runesB := []rune(b)
	i := len(runeseA) - 1
	j := len(runesB) - 1
	for i >= 0 && j >= 0 && runeseA[i] == runesB[j] {
		i--
		j--
	}
	return string(runeseA[i+1:])
}
