package exer8
import (
	"fmt"
	"math"
)
// TODO: The Point struct, NewPoint function, .String and .Norm methods

type Point struct{
	x float64
	y float64
}

func NewPoint(x float64, y float64)  Point{
	return Point{x,y}
}

func (pt Point) String() string{
	return fmt.Sprintf("(%v, %v)", pt.x, pt.y); 
}

func (pt Point) Norm() float64{
	var norm float64 = pt.x*pt.x + pt.y*pt.y 
	norm = math.Sqrt(norm)
	
	return norm
}