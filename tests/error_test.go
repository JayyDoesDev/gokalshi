package gokalshi

import (
	"testing"

	"github.com/jayydoesdev/gokalshi"
)

func TestError(t *testing.T) {
	got, _ := gokalshi.SprintfWithError(123, false)
	want := "123"
	if got != want {
		t.Errorf("expected %v, got %v", want, got)
	}
}
