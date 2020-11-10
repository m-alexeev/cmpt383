package main

import (
	"fmt"
	"../exer9"
	"math"
)


func main()  {

	pt := exer9.Point{1, 2}
	fmt.Println(pt)
	pt.Scale(5)
	fmt.Println(pt)

	p := exer9.Point{1, 0}
	p.Rotate(math.Pi / 2)
	fmt.Println(p)
	p.Rotate(math.Pi / 2)
	fmt.Println(p)



	a:= exer9.RandomArray(10, 100)
	fmt.Println(a)

	arr := []int{1,2,3,4,5,6,7,8,9,10}
	
	fmt.Println(exer9.MeanStddev(arr,5))

	arr2 := []float64{3,4,1,5,6,10,2,3,5,2}
	exer9.InsertionSort(arr2)
	arr3 := []float64{3,4,1,5,6,10,2,3,5,2}
	exer9.QuickSort(arr3)
}