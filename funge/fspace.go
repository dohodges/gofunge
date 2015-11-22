package funge

import (
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type point string

func makePoint(v Vector) point {
	s := make([]string, v.Size())
	for i := 0; i < v.Size(); i++ {
		value := v.Get(Axis(i))
		s[i] = strconv.FormatInt(int64(value), 36)
	}
	return point(strings.Join(s, ","))
}

func (p point) Vector() Vector {
	s := strings.Split(string(p), ",")
	v := NewVector(len(s))
	for i, str := range s {
		value, err := strconv.ParseInt(str, 36, 32)
		if err != nil {
			panic("gofunge.point: invalid point")
		}
		v.Set(Axis(i), int32(value))
	}
	return v
}

type FungeSpace struct {
	funge Funge
	data  map[point]rune
}

func NewFungeSpace(funge Funge) *FungeSpace {
	return &FungeSpace{
		funge: funge,
		data:  make(map[point]rune),
	}
}

func (fs *FungeSpace) Load(reader io.Reader) error {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	address := fs.funge.Origin()

	for _, r := range bytes.Runes(b) {
		switch r {
		case 012, 015:
			if fs.funge > 1 {
				address = address.Add(fs.funge.Delta(YAxis, Forward))
				address.Set(XAxis, 0)
			}
		case 014:
			if fs.funge > 2 {
				address = address.Add(fs.funge.Delta(ZAxis, Forward))
				address.Set(XAxis, 0)
				address.Set(YAxis, 0)
			} else if fs.funge == 2 {
				address = address.Add(fs.funge.Delta(YAxis, Forward))
				address.Set(XAxis, 0)
			}
		default:
			fs.Put(address, r)
			address = address.Add(fs.funge.Delta(XAxis, Forward))
		}
	}

	return nil
}

func (fs *FungeSpace) Get(address Vector) rune {
	if r, exists := fs.data[makePoint(address)]; exists {
		return r
	} else {
		return ' '
	}
}

func (fs *FungeSpace) Put(address Vector, r rune) {
	fs.data[makePoint(address)] = r
}

func (fs *FungeSpace) Clear() {
	fs.data = make(map[point]rune)
}
