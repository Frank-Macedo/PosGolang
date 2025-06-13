package math

type math struct {
	a int
	b int
}

func newMath(a, b int) *math {
	return &math{
		a: a,
		b: b,
	}
}
