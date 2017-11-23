package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var pop uint64
	for i := uint(0); i < 8; i++ {
		pop += (x >> i) & 1
	}
	return int(pop)
}
