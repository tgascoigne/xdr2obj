package types

import (
	"math"
)

type Ptr32 uint32

func (p Ptr32) Valid() bool {
	return p != 0
}

type Float16 uint16

func (i Float16) Value() float32 {
	/* Lovingly adapted from http://stackoverflow.com/a/15118210 */
	t1 := uint32(i & 0x7fff)
	t2 := uint32(i & 0x8000)
	t3 := uint32(i & 0x7c00)
	t1 <<= 13
	t2 <<= 16
	t1 += 0x38000000
	if t3 == 0 {
		t1 = 0
	}
	t1 |= t2
	return math.Float32frombits(t1)
}

type UV struct {
	U Float16
	V Float16
}

type Vec2 struct {
	X float32
	Y float32
}

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

type Vec3h struct {
	X Float16
	Y Float16
	Z Float16
}

type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

type WorldCoord Vec3

type WorldCoordh Vec3h

type Tri struct {
	A uint16
	B uint16
	C uint16
}
