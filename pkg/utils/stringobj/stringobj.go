package stringobj

import "fmt"

func Map2list(m map[string]string) string {
	s := ""
	for k, v := range m {
		s = s + fmt.Sprintf("%s:%s ", k, v)
	}
	return s
}
