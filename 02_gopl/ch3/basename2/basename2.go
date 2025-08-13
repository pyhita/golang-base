package basename2

import "strings"

func basename(s string) string {
	i := strings.LastIndex(s, "/")
	s = s[i+1:]

	if j := strings.LastIndex(s, "."); j >= 0 {
		s = s[:j]
	}

	return s
}
