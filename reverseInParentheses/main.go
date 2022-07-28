package main

import "strings"

func main() {

}

func reverseInParentheses(inputString string) string {
	s := strings.Split(inputString, "(")

	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], ")") {
			part := strings.Split(s[i], ")")
			part[0] = reverseText(part[0])
			s[i] = part[0] + part[1]
		}
	}

	return strings.Join(s, "")
}

// func reverseInParentheses(inputString string) string {

// 	stringSplitted := strings.Split(inputString, "(")
// 	endSplitted := strings.Split(stringSplitted[1], ")")
// 	reversedCode := reverseText(endSplitted[0])

// 	return stringSplitted[0] + reversedCode + endSplitted[1]
// }

func reverseText(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
