package lib

func SetBit(bitmap []byte, pos int) {
	index := pos / 8
	offset := uint(pos % 8)
	bitmap[index] |= 1 << offset
}

func ClearBit(bitmap []byte, pos int) {
	index := pos / 8
	offset := uint(pos % 8)
	bitmap[index] &= ^(1 << offset)
}

func GetBit(bitmap []byte, pos int) int {
	index := pos / 8
	offset := uint(pos % 8)
	bit := (bitmap[index] >> offset) & 1
	return int(bit)
}

func ToggleBit(bitmap []byte, pos int) {
	index := pos / 8
	offset := uint(pos % 8)
	bitmap[index] ^= 1 << offset
}
