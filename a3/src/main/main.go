package main

import(
	"../blockchain"
	"fmt"
)


func main(){
	b0 := blockchain.Initial(16)
	fmt.Println(b0.HashString())
	b1 := b0.Next("Message")
	fmt.Println(b1.HashString())

}

