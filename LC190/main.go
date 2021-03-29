package main

func reverseBits(n uint32) (rev uint32) {
	for i := 0; i < 32 && n > 0; i++ {
		rev |= n & 1 << (31 - i)
		n >>= 1
	}
	return
}

func main() {
	var n uint32
	n = 43261596
	reverseBits(n)
}
