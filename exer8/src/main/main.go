package main

import(
	"fmt"
	"../exer8"
)


func main(){
	
	fmt.Println(exer8.Hailstone(5))
	fmt.Println(exer8.HailstoneSequenceAppend(5))
	fmt.Println(exer8.HailstoneSequenceAllocate(5))

	pt := exer8.NewPoint(3, 4.5)
	fmt.Println(pt) // should print (3, 4.5)
	fmt.Println(pt.String() == "(3, 4.5)") // should print true

	pt1 := exer8.NewPoint(3, 4)
	fmt.Println(pt1.Norm() == 5.0)
}