package day17

type Position struct {
	i, j int
}

func (p Position) left() Position {
	p.j--
	return p
}

func (p Position) right() Position {
	p.j++
	return p
}

func (p Position) top() Position {
	p.i--
	return p
}

func (p Position) bottom() Position {
	p.i++
	return p
}

func (p Position) direction(dest Position) Direction {

	if p.i == dest.i {
		if p.j > dest.j {
			return directionLeft
		} else if p.j < dest.j {
			return directionRight
		} else {
			panic("same")
		}
	}

	if p.j == dest.j {
		if p.i > dest.i {
			return directionTop
		} else if p.i < dest.i {
			return directionBottom
		} else {
			panic("same")
		}
	}

	panic("same")
}

func (p Position) neighbours() []Position {
	return []Position{
		p.left(),
		p.right(),
		p.top(),
		p.bottom(),
	}
}
