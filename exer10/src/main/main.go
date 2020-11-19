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


	pt:= exer10.Point{X:3,Y:4}
	exer10.TurnDouble(&pt,3*math.Pi/2)
	fmt.Println(pt.String())

	tri := exer10.Triangle{A:exer10.Point{X:1,Y:2}, B:exer10.Point{X:-3,Y:4}, C:exer10.Point{X:5,Y:-6}}
	exer10.TurnDouble(&tri, math.Pi)
	fmt.Println(tri.String())
}