package dice

import "math/rand"


type PercentilePool struct {
	*DicePool
}

func NewPercentilePool(sides, number int) *DicePool {
	return &PercentilePool{ &DicePool{ sides, number } }
}

func (d *PercentilePool) Roll() {
	if d.number > 0 {
		if cap(d.Values) == 0 {
			d.Values = make([]int, 0, number)
		}
		for n := d.number; n > 0; n-- {
			d.Values[l] = d.rand() * d.rand()
		}
	}
}