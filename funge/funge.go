package funge

import (
	"math/rand"
	"time"
)

type Funge int
type Axis int
type Direction int

const (
	Unefunge Funge = iota + 1
	Befunge
	Trefunge
)

const (
	XAxis Axis = iota
	YAxis
	ZAxis
)

const (
	Forward Direction = iota
	Backward
)

func init() {
	rand.Seed(time.Now().Unix())
}

func (f Funge) Origin() Vector {
	return NewVector(int(f))
}

func (f Funge) OriginDelta() Vector {
	return f.Delta(XAxis, Forward)
}

func (f Funge) Delta(axis Axis, direction Direction) Vector {
	delta := NewVector(int(f))
	if direction == Forward {
		delta.Set(axis, 1)
	} else {
		delta.Set(axis, -1)
	}
	return delta
}

func (f Funge) RandomAxis() Axis {
	return Axis(rand.Intn(int(f)))
}

func (f Funge) RandomDirection() Direction {
	return Direction(rand.Intn(2))
}

func (f Funge) LeftTurnTransform() *Matrix {
	transform := IdentityMatrix(int(f))
	transform.Set(0, 0, 0)
	transform.Set(0, 1, -1)
	transform.Set(1, 0, 1)
	transform.Set(1, 1, 0)
	return transform
}

func (f Funge) RightTurnTransform() *Matrix {
	transform := IdentityMatrix(int(f))
	transform.Set(0, 0, 0)
	transform.Set(0, 1, 1)
	transform.Set(1, 0, -1)
	transform.Set(1, 1, 0)
	return transform
}
