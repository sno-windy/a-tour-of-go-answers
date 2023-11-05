package main

import "golang.org/x/tour/wc"

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	words := make([]string, 0)
	var start, end, wordLen int
	for i, rune := range s {
		wordLen += 1
		if string(rune) != " " {
			continue
		}
		end = i
		words = append(words, s[start:end])
		start = i+1
		wordLen = 0
	}
	if wordLen > 0 {
		words = append(words, s[start:len(s)])
	}	
	for _, v := range words {
		_, ok := result[v]
		if ok {
			result[v] += 1
		} else {
			result[v] = 1
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
