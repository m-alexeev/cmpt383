package main

import(
	"../blockchain"
	"fmt"
)


func main(){
	b0 := blockchain.Initial(16)
	fmt.Println(b0)
}

