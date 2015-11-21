package funge

type Pointer struct {
	funge   Funge
	address Vector
	delta   Vector
}

func NewPointer(funge Funge) *Pointer {
	return &Pointer{
		funge:   funge,
		address: funge.Origin(),
		delta:   funge.OriginDelta(),
	}
}

func (p *Pointer) Address() Vector {
	return p.address
}

func (p *Pointer) Delta() Vector {
	return p.delta
}

func (p *Pointer) East() {
	p.Go(p.funge.Delta(XAxis, Forward))
}

func (p *Pointer) West() {
	p.Go(p.funge.Delta(XAxis, Backward))
}

func (p *Pointer) South() {
	p.Go(p.funge.Delta(YAxis, Forward))
}

func (p *Pointer) North() {
	p.Go(p.funge.Delta(YAxis, Backward))
}

func (p *Pointer) Up() {
	p.Go(p.funge.Delta(ZAxis, Forward))
}

func (p *Pointer) Down() {
	p.Go(p.funge.Delta(ZAxis, Backward))
}

func (p *Pointer) Away() {
	axis := p.funge.RandomAxis()
	direction := p.funge.RandomDirection()
	p.Go(p.funge.Delta(axis, direction))
}

func (p *Pointer) TurnLeft() {
	if p.funge >= 2 {
		p.delta = p.delta.Transform(p.funge.LeftTurnTransform())
	}
}

func (p *Pointer) TurnRight() {
	if p.funge >= 2 {
		p.delta = p.delta.Transform(p.funge.RightTurnTransform())
	}
}

func (p *Pointer) Go(delta Vector) {
	p.delta = delta
}

func (p *Pointer) Next(steps int) {
	for s := 0; s < steps; s++ {
		p.address = p.address.Add(p.delta)
	}
}
