package string

import (
	"strconv"
	"strings"
	"testing"
)

func concatenation(entity map[string][]string) string {
	var headers string
	for key, values := range entity {
		for _, value := range values {
			if headers == "" {
				headers = key + ": " + value
			} else {
				headers = headers + "; " + key + ": " + value
			}
		}
	}
	return headers
}

func stringBuilder(entity map[string][]string) string {
	var builder strings.Builder
	for key, values := range entity {
		for _, value := range values {
			if builder.Len() > 0 {
				builder.WriteString("; ")
			}
			builder.WriteString(key)
			builder.WriteString(": ")
			builder.WriteString(value)
		}
	}
	return builder.String()
}

const mockValuesCount = 50

var mockEntity = func() map[string][]string {
	var entitiesMap = make(map[string][]string)
	for i := 0; i < mockValuesCount; i++ {
		var iStr = strconv.Itoa(i)
		key := "key" + iStr
		value1 := "value" + iStr
		value2 := "value" + iStr
		entitiesMap[key] = []string{value1, value2}
	}
	return entitiesMap
}()

func BenchmarkKibana(b *testing.B) {
	b.Run("concatenation", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			concatenation(mockEntity)
		}
	})

	b.Run("stringBuilder", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			stringBuilder(mockEntity)
		}
	})
}
