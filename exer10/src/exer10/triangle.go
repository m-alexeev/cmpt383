package exer10

// TODO: Point (with everything from exercise 8) and methods that modify them in-place
import (
	"fmt"
	"math"
)

type Point struct{
	X, Y float64
}


type Triangle struct{
	A, B, C Point
}


func (t Triangle) String() string{
	return fmt.Sprintf("[%s %s %s]", t.A, t.B, t.C)
}


func (pt Point) String() string{
	return fmt.Sprintf("(%v, %v)", pt.X, pt.Y); 
}

func (pt Point) Norm() float64{
	var norm float64 = pt.X*pt.X + pt.Y*pt.Y 
	norm = math.Sqrt(norm)
	
	return norm
}

func (pt *Point) Scale(f float64) {
	pt.X = pt.X * f; 
	pt.Y = pt.Y * f;
}

func (pt *Point) Rotate(a float64){
	y:= pt.X * math.Sin(a) + pt.Y * math.Cos(a)
	pt.X = pt.X * math.Cos(a) - pt.Y * math.Sin(a)
	pt.Y = y
}