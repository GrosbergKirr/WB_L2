package num4

import (
	"strings"
)

func Equal(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	exists := make(map[rune]struct{})
	for _, v := range str1 {
		exists[v] = struct{}{}
	}
	for _, v := range str2 {
		if _, ok := exists[v]; !ok {
			return false
		}
	}

	return true
}

func Anagram(data []string) map[string][]string {
	res := make(map[string][]string)
	exists := make(map[string]struct{})

	for i, v := range data {
		exists[v] = struct{}{}
		for j := i + 1; j < len(data); j++ {
			if _, ok := exists[data[j]]; ok {
				continue
			}

			if Equal(v, data[j]) {
				res[v] = append(res[v], data[j])
				exists[data[j]] = struct{}{}
			}
		}
	}

	return res
}

func LowCase(data []string) []string {
	res := make([]string, 0, len(data))

	for _, v := range data {
		str := strings.ToLower(v)
		res = append(res, str)
	}

	return res
}
