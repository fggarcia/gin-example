package handle

import (
	"strconv"
	"testing"
	"unique"
)

func BenchmarkHandle(b *testing.B) {
	var handleMap = make(map[int]unique.Handle[string], 10)
	var stringMap = make(map[int]string, 10)
	b.Run("handle", func(b *testing.B) {
		var imod int
		var h unique.Handle[string]

		for i := 0; i < 10; i++ {
			handleMap[i] = unique.Make(buildStr(i))
		}

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			imod = i % 10
			h = unique.Make(buildStr(imod))
			consumeBool(handleMap[imod] == h)
		}
	})
	b.Run("string", func(b *testing.B) {
		var imod int
		var str string

		for i := 0; i < 10; i++ {
			stringMap[i] = buildStr(i)
		}

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			imod = i % 10
			str = buildStr(imod)
			consumeBool(stringMap[imod] == str)
		}
	})
}

func consumeBool(b bool) bool {
	return b
}

func buildStr(i int) string {
	return "string" + strconv.Itoa(i)
}
