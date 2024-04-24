package num1

import (
	"fmt"
	"strconv"
	"strings"
)

func UnpackString(s string) (string, error) {

	var result strings.Builder
	var countStr strings.Builder

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			countStr.WriteByte(s[i])
		} else {
			if countStr.Len() > 0 {
				count, err := strconv.Atoi(countStr.String())
				if err != nil {
					return "", fmt.Errorf("некорректная строка: %v", err)
				}
				result.WriteString(strings.Repeat(string(s[i-1]), count))
				countStr.Reset()
			}
		}
	}

	if countStr.Len() > 0 {
		return "", fmt.Errorf("некорректная строка")
	}

	return result.String(), nil
}
