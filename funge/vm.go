package funge

import (
	"io"
)

type VirtualMachine struct {
	funge  Funge
	fspace *FungeSpace
	stack  *Stack
	ip     *Pointer
	term   *Terminal
}

func NewVirtualMachine(funge Funge) *VirtualMachine {
	return &VirtualMachine{
		funge:  funge,
		fspace: NewFungeSpace(funge),
		stack:  NewStack(funge),
		ip:     NewPointer(funge),
		term:   NewTerminal(),
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
		vm.next()
		for v := vm.fetch(); v != '"'; v = vm.fetch() {
			vm.stack.Push(v)
			vm.next()
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
		vm.term.OutputDecimal(vm.stack.Pop())
	case ',':
		vm.term.OutputCharacter(vm.stack.Pop())
	case '#':
		vm.next()
	case 'p':
		addr := vm.popStorageAddress()
		value := vm.stack.Pop()
		vm.fspace.Put(addr, value)
	case 'g':
		addr := vm.popStorageAddress()
		value := vm.fspace.Get(addr)
		vm.stack.Push(value)
	case '&':
		d := vm.term.InputDecimal()
		vm.stack.Push(d)
	case '~':
		c := vm.term.InputCharacter()
		vm.stack.Push(c)
	case '@':
		return io.EOF
	default:
	}
	vm.next()

	return nil
}

func (vm *VirtualMachine) fetch() rune {
	return vm.fspace.Get(vm.ip.Address())
}

func (vm *VirtualMachine) next() {
	vm.ip.Next()

	// handle wrapping
	if !vm.fspace.InBounds(vm.ip.Address()) {
		// reverse the ip and backtrack to the other bound
		vm.ip.Reverse()
		for vm.ip.Next(); vm.fspace.InBounds(vm.ip.Address()); vm.ip.Next() {
		}

		// reverse again and proceed
		vm.ip.Reverse()
		vm.ip.Next()
	}
}

func (vm *VirtualMachine) popStorageAddress() Vector {
	addr := vm.stack.PopVector()
	return addr.Add(vm.ip.StorageOffset())
}
