package dice


type Selectable interface {
	Select(func(int) bool)
}

func AtLeast(s Selectable, target int) int {
	return s.Select(func(x int) bool {
		return x >= target
	})
}

func NoMoreThan(s Selectable, target int) int {
	return s.Select(func(x int) bool {
		return x <= target
	})
}

func SameAs(s Selectable, target int) int {
	return s.Select(func(x int) bool {
		return x <= target
	})
}