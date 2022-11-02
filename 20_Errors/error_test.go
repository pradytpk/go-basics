package main

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"19:00:01", true},
		{"1:-3:44", false},
		{"", false},
		{"aa:bb:cc", false},
	}
	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v:%v,error should be nil", data.time, err)
		}
	}
}
