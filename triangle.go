package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcArea() float64 {
	s := (t.Side + t.Side + t.Side) / 2
	return math.Sqrt(s * (s - t.Side) * (s - t.Side) * (s - t.Side))
}

func (t Triangle) CalcPerimeter() float64 {
	return 3 * t.Side
}
