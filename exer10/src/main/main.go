package main


import (
	"../exer10"
	"fmt"
	"math"
)

func main(){

	// for i:=uint(0); i <30 ; i ++{
	// 	fmt.Println(exer10.Fibbonaci(i))
	// }


	pt:= exer10.Point{3,4}
	exer10.TurnDouble(&pt,3*math.Pi/2)
	fmt.Println(pt.String())

	tri := exer10.Triangle{exer10.Point{1,2}, exer10.Point{-3,4}, exer10.Point{5,-6}}
	exer10.TurnDouble(&tri, math.Pi)
	fmt.Println(tri.String())
}