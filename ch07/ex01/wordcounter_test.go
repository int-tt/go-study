package wordcounter

import "testing"

func TestWordCounter(t *testing.T) {
	date := map[string]WordCounter{
		"hello":     1,
		"test":      1,
		"1\n2\n3\n": 3,
	}
	for input, expeted := range date {
		var c WordCounter
		c.Write([]byte(input))
		if c != expeted {
			t.Errorf("'%s' shoud be %d characters\nanswer:%d\n", input, expeted, c)
		}
	}

}
