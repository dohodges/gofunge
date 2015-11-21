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

	fmt.Println(vm.fspace)

	for {
		if err := vm.cycle(); err != nil {
			return err
		}
	}

	return nil
}

func (vm *VirtualMachine) cycle() error {
	instr := vm.fetch()

	switch instr {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		vm.stack.Push(instr)
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
	case '.':
		value, err := vm.stack.Pop()
		if err != nil {
			return err
		}
		fmt.Print(value)
	default:
	}
	vm.ip.Next(1)

	return nil
}

func (vm *VirtualMachine) fetch() rune {
	if instr, ok := vm.fspace.Get(vm.ip.Address()); ok {
		return instr
	} else {
		return ' '
	}
}
