package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var pcForLoop [256]byte
	for i := range pcForLoop {
		pcForLoop[i] = pcForLoop[i/2] + byte(i&1)
	}
	return int(pcForLoop[byte(x>>(0*8))] +
		pcForLoop[byte(x>>(1*8))] +
		pcForLoop[byte(x>>(2*8))] +
		pcForLoop[byte(x>>(3*8))] +
		pcForLoop[byte(x>>(4*8))] +
		pcForLoop[byte(x>>(5*8))] +
		pcForLoop[byte(x>>(6*8))] +
		pcForLoop[byte(x>>(7*8))])
}
