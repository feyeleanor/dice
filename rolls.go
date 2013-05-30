package dice


func D4(n int) DicePool {
	return NewDicePool(4, n)
}

func D6(n int) DicePool {
	return NewDicePool(6, n)
}

func D8(n int) DicePool {
	return NewDicePool(8, n)
}

func D10(n int) DicePool {
	return NewDicePool(10, n)
}

func D12(n int) DicePool {
	return NewDicePool(12, n)
}

func D20(n int) DicePool {
	return NewDicePool(20, n)
}

func D66(n int) PercentilePool {
	return NewPercentilePool(6, n)
}

func D100(n int) PercentilePool {
	return NewPercentilePool(10, n)
}