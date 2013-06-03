package dice


type PercentilePool struct {
	*DicePool
}

func NewPercentilePool(sides, number int) *PercentilePool {
	return &PercentilePool{ NewDicePool(sides, number) }
}

func (d *PercentilePool) Roll() {
	if d.number > 0 {
		if cap(d.values) == 0 {
			d.values = make([]int, 0, d.number)
		}
		for n := d.number - 1; n > -1; n-- {
			d.values[n] = d.rand() * d.rand()
		}
	}
}