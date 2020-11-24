package main

import(
	"../blockchain"
	// "encoding/hex"

	"fmt"
)


func main(){

	
	b0 := blockchain.Initial(7)
	b0.MineSequential()
	Res:= b0.Proof
	for i:= uint64(0); i < Res; i++{
		b0.SetProof(i)
		if (b0.ValidHash()){
			fmt.Println(b0.Proof,b0.ValidHash())
		}
	}
	b1 := b0.Next("this is an interesting message")
	for i:= uint64(0); i < 400; i++{
		b1.SetProof(i)
		if (b1.ValidHash()){
			fmt.Println(b1.Proof,b1.ValidHash())
		}
	}
	// b0.Mine(1)
	// fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	// b1 := b0.Next("this is an interesting message")
	// b1.Mine(1)
	// fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	// b2 := b1.Next("this is not interesting")
	// b2.Mine(1)
	// fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))

}

