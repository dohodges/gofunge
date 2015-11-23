package funge

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Terminal struct {
	stdin io.RuneReader
}

func NewTerminal() *Terminal {
	return &Terminal{
		stdin: bufio.NewReader(os.Stdin),
	}
}

func (t *Terminal) OutputDecimal(value rune) {
	fmt.Printf(`%d `, value)
}

func (t *Terminal) OutputCharacter(value rune) {
	fmt.Printf(`%s`, string(value))
}

func (t *Terminal) InputDecimal() int32 {
	seq := make([]rune, 0)
	inNumericSeq := false

	for {
		r := t.readRune()
		if !inNumericSeq && (r >= '0' && r <= '9') {
			inNumericSeq = true
		} else if inNumericSeq && (r < '0' || r > '9') {
			break
		}

		if inNumericSeq {
			seq = append(seq, r)
		}
	}

	for {
		if value, err := strconv.ParseInt(string(seq), 10, 32); err == nil {
			return int32(value)
		} else if err == strconv.ErrRange {
			// overflow, trim last digit from seq
			seq = seq[:len(seq)-1]
		} else {
			panic(err)
		}
	}
}

func (t *Terminal) InputCharacter() rune {
	return t.readRune()
}

func (t *Terminal) readRune() rune {
	r, _, err := t.stdin.ReadRune()
	if err != nil {
		panic("funge.Terminal: invalid rune")
	}

	return r
}
