package main


import (
	"../exer10"
	"fmt"
)

func main(){

	for i:=uint(0); i <30 ; i ++{
		fmt.Println(exer10.Fibbonaci(i))
	}
}