package main

import (
	"testing"
)

func TestConnma(t *testing.T) {
	data := map[string]string{
		"123":       "123",
		"12345":     "12,345",
		"123456789": "123,456,789",
	}
	for input, expeted := range data {
		result := comma(input)
		if result != expeted {
			t.Errorf("\ninput:\t%s\nresult:\t%s\nexpeted:%s", input, result, expeted)
		}
	}
}
