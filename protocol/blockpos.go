package protocol

type BlockPos struct {
	x, y, z int64
}

func FromLong(serialized int64) BlockPos {
	x := serialized >> 38
	y := (serialized >> 26) & 0xFFF
	z := serialized << 38 >> 38
	if x >= 2^25 {
		x -= 2 ^ 26
	}
	if y >= 2^11 {
		y -= 2 ^ 12
	}
	if z >= 2^25 {
		z -= 2 ^ 26
	}
	return BlockPos{x: x, y: y, z: z}
}
