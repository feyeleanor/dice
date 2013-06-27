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
	return rand.Intn(d.sides) + 1
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
	return
}

func (d *DicePool) Roll() {
	if d.number > 0 {
		if cap(d.values) == 0 {
			d.values = make([]int, d.number)
		}
		for n := d.number - 1; n > -1; n-- {
			d.values[n] = d.rand()
		}
	}
}
