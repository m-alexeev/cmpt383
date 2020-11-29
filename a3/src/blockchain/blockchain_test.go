package blockchain

import (
	// "github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleMining(t *testing.T){
	
	//Create 3 Sequential Blocks
	b0s :=  Initial(5)
	b0s.MineSequential()

	b1s := b0s.Next("Some Hash") 
	b1s.MineSequential() 

	b2s := b1s.Next("Another Hash")
	b2s.MineSequential()


	//Create 3 Blocks to mine Concurrently
	b0c := Initial(5)
	b0c.Mine(1)

	b1c := b0c.Next("Some Hash")
	b1c.Mine(1)

	b2c := b1c.Next("Another Hash")
	b2c.Mine(1)

	//Compare the valid Proofs and check for proof validity
	if b0c.Proof != b0s.Proof {
		t.Error("Block 0 Sequential and Concurrent Proofs are different")
	}	
	if !b0c.ValidHash(){
		t.Error("B1 Hash is Invalid")
	}
	if b1c.Proof != b1s.Proof{
		t.Error("Block 1 Sequential and Concurrent Proofs are different")
	}
	if !b1c.ValidHash(){
		t.Error("B1 Hash is Invalid")
	}
	if b2c.Proof != b2s.Proof{
		t.Error("Block 2 Sequential and Concurrent Proofs are different")
	}
	if !b2c.ValidHash(){
		t.Error("B2 Hash is Invalid")
	}

}

// TODO: some useful tests of Blocks
