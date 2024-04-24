package patterns

// Circle represents a circle shape.
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

// Square represents a square shape.
type Square struct {
	SideLength float64
}

func (s *Square) Accept(visitor Visitor) {
	visitor.VisitSquare(s)
}

// Visitor represents the interface for visitors.
type Visitor interface {
	VisitCircle(circle *Circle)
	VisitSquare(square *Square)
}

//Next, we implement concrete visitor classes to calculate the area and perimeter:

// AreaVisitor calculates the area of shapes.
type AreaVisitor struct {
	AreaTotal float64
}

func (v *AreaVisitor) VisitCircle(circle *Circle) {
	v.AreaTotal += 3.14 * circle.Radius * circle.Radius
}

func (v *AreaVisitor) VisitSquare(square *Square) {
	v.AreaTotal += square.SideLength * square.SideLength
}

// PerimeterVisitor calculates the perimeter of shapes.
type PerimeterVisitor struct {
	PerimeterTotal float64
}

func (v *PerimeterVisitor) VisitCircle(circle *Circle) {
	v.PerimeterTotal += 2 * 3.14 * circle.Radius
}

func (v *PerimeterVisitor) VisitSquare(square *Square) {
	v.PerimeterTotal += 4 * square.SideLength
}
