package funge

import (
	"fmt"
	"io"
)

type VirtualMachine struct {
	funge  Funge
	fspace *FungeSpace
	stack  *Stack
	ip     *Pointer
}

func NewVirtualMachine(funge Funge) *VirtualMachine {
	return &VirtualMachine{
		funge:  funge,
		fspace: NewFungeSpace(funge),
		stack:  NewStack(),
		ip:     NewPointer(funge),
	}
}

func (vm *VirtualMachine) Reset() {
	vm.fspace.Clear()
	vm.stack.Clear()
	vm.ip = NewPointer(vm.funge)
}

func (vm *VirtualMachine) LoadProgram(reader io.Reader) error {
	return vm.fspace.Load(reader)
}

func (vm *VirtualMachine) Run() error {

	for {
		if err := vm.step(); err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
	}

	return nil
}

func (vm *VirtualMachine) step() error {
	instr := vm.fetch()

	switch instr {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		vm.stack.Push(instr - 48)
	case '+':
		a, b := vm.stack.Pop(), vm.stack.Pop()
		vm.stack.Push(a + b)
	case '-':
		a, b := vm.stack.Pop(), vm.stack.Pop()
		vm.stack.Push(b - a)
	case '*':
		a, b := vm.stack.Pop(), vm.stack.Pop()
		vm.stack.Push(a * b)
	case '/':
		a, b := vm.stack.Pop(), vm.stack.Pop()
		if a == 0 {
			vm.stack.Push(0) // 93-rule: if a is zero, ask the user what result they want
		} else {
			vm.stack.Push(b / a)
		}
	case '%':
		a, b := vm.stack.Pop(), vm.stack.Pop()
		vm.stack.Push(b % a)
	case '!':
		if a := vm.stack.Pop(); a == 0 {
			vm.stack.Push(1)
		} else {
			vm.stack.Push(0)
		}
	case '`':
		if a, b := vm.stack.Pop(), vm.stack.Pop(); b > a {
			vm.stack.Push(1)
		} else {
			vm.stack.Push(0)
		}
	case '^':
		vm.ip.North()
	case '>':
		vm.ip.East()
	case 'v':
		vm.ip.South()
	case '<':
		vm.ip.West()
	case '?':
		vm.ip.Away()
	case '_':
		if a := vm.stack.Pop(); a == 0 {
			vm.ip.East()
		} else {
			vm.ip.West()
		}
	case '|':
		if a := vm.stack.Pop(); a == 0 {
			vm.ip.South()
		} else {
			vm.ip.North()
		}
	case '"':
		vm.ip.Next()
		for v := vm.fetch(); v != '"'; v = vm.fetch() {
			vm.stack.Push(v)
			vm.ip.Next()
		}
	case ':':
		a := vm.stack.Pop()
		vm.stack.Push(a, a)
	case '\\':
		a, b := vm.stack.Pop(), vm.stack.Pop()
		vm.stack.Push(a, b)
	case '$':
		vm.stack.Pop()
	case '.':
		value := vm.stack.Pop()
		fmt.Printf(`%d `, value)
	case ',':
		value := vm.stack.Pop()
		fmt.Printf(`%s`, string(value))
	case '#':
		vm.ip.Next()
	case '@':
		return io.EOF
	default:
	}
	vm.ip.Next()

	return nil
}

func (vm *VirtualMachine) fetch() rune {
	if instr, ok := vm.fspace.Get(vm.ip.Address()); ok {
		return instr
	} else {
		return ' '
	}
}
