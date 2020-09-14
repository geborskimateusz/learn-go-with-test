package structsmethodsandinterfaces

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(circle Circle) float64 {
	return 3.14 * circle.Radius * circle.Radius
}
