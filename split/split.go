package split

import "strings"

func splitByDashUsingSplit(s string, splitter byte) (string, string) {
	parts := strings.Split(s, string(splitter))
	if len(parts) != 2 {
		return "", ""
	}
	return parts[0], parts[1]
}

// Función optimizada usando slices de bytes
func splitByDashUsingBytes(s string, splitter byte) (string, string) {
	data := []byte(s)
	var start, end int

	for i, b := range data {
		if b == splitter {
			start = i
			break
		}
	}

	if start > 0 {
		end = start + 1
	}
	return string(data[:start]), string(data[end:])
}
