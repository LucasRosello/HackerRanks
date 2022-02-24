package main

func commonCharacterCount(s1 string, s2 string) int {
	characters1 := countLetters(s1)
	characters2 := countLetters(s2)

	repeatedCharacters := countRepeatedCharacters(characters1, characters2)

	return repeatedCharacters
}

func countLetters(s string) map[string]int {
	letters := make(map[string]int)
	for _, v := range s {
		letters[string(v)]++
	}

	return letters
}

func countRepeatedCharacters(c1 map[string]int, c2 map[string]int) int {
	var count int

	for key := range c1 {
		if c1[key] < c2[key] {
			count += c1[key]
		} else {
			count += c2[key]
		}
	}

	return count
}
