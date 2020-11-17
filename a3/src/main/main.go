package main

import(
	"../blockchain"
	"encoding/hex"

	"fmt"
)


func main(){
	b0 := blockchain.Initial(19)
	b0.SetProof(56231)
	fmt.Println(b0.ValidHash())

	b1 := b0.Next("hash example 1234")
	b1.SetProof(346082)
	fmt.Println(b1.ValidHash())


	fmt.Println(hex.EncodeToString(b0.CalcHash()))
}

