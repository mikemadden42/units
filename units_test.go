package main

import (
	"math"
	"testing"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		in   uint64
		want string
	}{
		{0, "0B"},
		{1, "1B"},
		{1023, "1023B"},
		{1024, "1K"},
		{1536, "1.5K"},
		{1 << 20, "1M"},
		{1 << 30, "1G"},
		{math.MaxUint64, "16E"},
	}
	for _, tt := range tests {
		if got := format(tt.in); got != tt.want {
			t.Errorf("format(%d) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
