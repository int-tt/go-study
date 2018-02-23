package bytecounter

import "testing"

func TestByteCounterWrite(t *testing.T) {
	date := map[string]ByteCounter{
		"hello": 5,
		"test":  4,
	}
	for input, expeted := range date {
		var c ByteCounter
		c.Write([]byte(input))
		if c != expeted {
			t.Errorf("'%s' shoud be %d characters\nanswer:%d\n", input, expeted, c)
		}
	}

}
