package decodestrings

import (
	"fmt"
	"strconv"
	"unicode"
)

func repeat(s string, n int) (ret string) {

	for i := 0; i < n; i++ {
		ret = fmt.Sprintf("%s%s", ret, s)
	}

	return
}

// DecodeString will decode a correctly formated enconded string
func DecodeString(s string) string {
	var ret string
	var repeatCount int
	var startEnclosure int
	var openBracketCount int

	for index, r := range s {

		switch {
		case unicode.IsLetter(r) && openBracketCount == 0:
			ret = fmt.Sprintf("%s%c", ret, r)
		case unicode.IsDigit(r) && openBracketCount == 0:
			nDigit, _ := strconv.Atoi(fmt.Sprintf("%c", r))
			fmted := fmt.Sprintf("%d%d", repeatCount, nDigit)
			repeatCount, _ = strconv.Atoi(fmted)
		case r == '[':
			if openBracketCount == 0 {
				startEnclosure = index
			}
			openBracketCount++
		case r == ']':
			openBracketCount--
			if openBracketCount == 0 {
				innerStr := DecodeString(s[startEnclosure+1 : index])
				ret = fmt.Sprintf("%s%s", ret, repeat(innerStr, repeatCount))
				repeatCount = 0
			}
		}

	}

	return ret
}
