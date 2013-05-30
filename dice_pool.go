package dice

import "math/rand"


type Filter	func(int) bool


type DicePool struct {
	sides	int
	number	int
	values	[]int
}

func NewDicePool(sides, number int) *DicePool {
	return &DicePool{
		sides:		sides,
		number:		number,
	}
}

func (d *DicePool) rand() int {
	return rand.Intn(d.sides - 1) + 1
}

func (d DicePool) Values() (r []int) {
	r = make([]int, len(d.values))
	copy(r, d.values)
	return
}

func (d DicePool) Sum() (r int) {
	for _, v := range d.values {
		r += v
	}
	return
}

func (d DicePool) Select(f Filter) (r int) {
	for _, v := range d.values {
		if f(v) {
			r++
		}
	}
}

func (d *DicePool) Roll() {
	if d.number > 0 {
		if cap(d.values) == 0 {
			d.values = make([]int, 0, number)
		}
		for n := d.number; n > 0; n-- {
			d.values[l] = d.rand()
		}
	}
}