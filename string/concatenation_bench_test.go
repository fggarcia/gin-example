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

		for b.Loop() {
			concatenation(mockEntity)
		}
	})

	b.Run("stringBuilder", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			stringBuilder(mockEntity)
		}
	})
}

func BenchmarkBuildKey(b *testing.B) {
	b.Run("fmt.Sprintf", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			consumeString(fmtSprintf(strconv.Itoa(i)))
		}
	})
	b.Run("slice", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			consumeString(sliceCompose(strconv.Itoa(i)))
		}
	})
	b.Run("concatStringBuilder", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			consumeString(concatStringBuilder(strconv.Itoa(i)))
		}
	})
	b.Run("simpleConcat", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			consumeString(simpleConcat(strconv.Itoa(i)))
		}
	})
}

func consumeString(str string) string {
	return str
}

func BenchmarkRegex(b *testing.B) {
	value := "value"
	b.Run("slice", func(b *testing.B) {
		b.ReportAllocs()
		cache := NewRegexCache(createRegexCompose)

		for b.Loop() {
			regex := cache.Get(value)
			consumeString(regex.String())
		}
	})
	b.Run("fmt.Sprintf", func(b *testing.B) {
		b.ReportAllocs()
		cache := NewRegexCache(createRegexFmt)

		for b.Loop() {
			regex := cache.Get(value)
			consumeString(regex.String())
		}
	})
}
