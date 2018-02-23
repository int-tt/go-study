package bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // int を ByteCounterへ変換
	return len(p), nil
}
