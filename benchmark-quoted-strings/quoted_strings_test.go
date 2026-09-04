package quotebench

import (
	"regexp"
	"strings"
	"testing"
)

const line = `MAT 1 "Material 1"`

var quoteRe = regexp.MustCompile(`"([^"]*)"`)

func extractRegex(s string) string {
	match := quoteRe.FindStringSubmatch(s)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func extractStrings(s string) string {
	first := strings.Index(s, `"`)
	last := strings.LastIndex(s, `"`)
	if first != -1 && last != -1 && first != last {
		return s[first+1 : last]
	}
	return ""
}

func BenchmarkRegex(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = extractRegex(line)
	}
}

func BenchmarkStrings(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = extractStrings(line)
	}
}