package test

import (
	"testing"

	appender "geborskimateusz.com/m/iteration"
)

func TestAppender(t *testing.T) {
	const given = "a"
	const expected = "aaaaa"

	result := appender.AppendViaAddAndAssignment(given)

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}

	result = appender.AppendViaSB(given)

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}

}

func BenchmarkAppendViaAddAndAssignment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appender.AppendViaAddAndAssignment("a")
	}
}
func BenchmarkAppendViaSB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appender.AppendViaSB("a")
	}
}
