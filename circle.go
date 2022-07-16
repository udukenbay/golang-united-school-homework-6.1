package golang_united_school_homework

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) CalcArea() float64 {
	return 3.142 * c.Radius * c.Radius
}

func (c Circle) CalcPerimeter() float64 {
	return 2 * 3.142 * c.Radius
}
