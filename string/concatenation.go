package string

import (
	"fmt"
	"regexp"
	"strings"
)

const key = "key"

func fmtSprintf(b string) string {
	return fmt.Sprintf("%s_%s", key, b)
}

func sliceCompose(b string) string {
	var s = make([]byte, 0, len(key)+len(b)+1)
	s = append(s, key...)
	s = append(s, '_')
	s = append(s, b...)
	return string(s)
}

func concatStringBuilder(b string) string {
	var sb strings.Builder
	//sb.Grow(len(key) + len(b) + 1)
	sb.WriteString(key)
	sb.WriteByte('_')
	sb.WriteString(b)
	return sb.String()
}

func simpleConcat(b string) string {
	return key + "_" + b
}

func createRegexFmt(key string) string {
	return fmt.Sprintf("%s_%s", "fallback_pcre_", key)
}

func createRegexCompose(key string) string {
	regex := make([]byte, 0, len("fallback_pcre_")+len(key)+1)
	regex = append(regex, []byte("fallback_pcre_")...)
	regex = append(regex, '_')
	regex = append(regex, []byte(key)...)
	return string(regex)
}

type RegexCache struct {
	cache map[string]*regexp.Regexp
	f     func(string) string
}

func NewRegexCache(f func(string) string) *RegexCache {
	return &RegexCache{
		cache: make(map[string]*regexp.Regexp),
		f:     f,
	}
}

func (rc *RegexCache) Get(pattern string) *regexp.Regexp {
	var result *regexp.Regexp
	var hit bool
	keyPattern := rc.f(pattern)
	if result, hit = rc.cache[keyPattern]; hit {
		return result
	} else {
		result, _ = regexp.Compile(keyPattern)
		rc.cache[keyPattern] = result
		return result
	}
}
