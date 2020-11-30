package blockchain

import (
	"testing"
	"time"
	"encoding/hex"
)


func TestInstructorValues(t *testing.T){
	b0 := Initial(7)
	b0.Mine(1)
	if b0.Proof != 385{
		t.Error("Block 0 Proof is not same as instructor proof")
	}
	if hex.EncodeToString(b0.Hash) != "379bf2fb1a558872f09442a45e300e72f00f03f2c6f4dd29971f67ea4f3d5300"{
		t.Error("Block 0 Hash is not the same as the instructor hash")
	}

	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	if b1.Proof != 20{
		t.Error("Block 1 Proof is not same as instructor proof")
	}
	if hex.EncodeToString(b1.Hash) != "4a1c722d8021346fa2f440d7f0bbaa585e632f68fd20fed812fc944613b92500"{
		t.Error("Block 1 Hash is not the same as the instructor hash")
	}

	b2 := b1.Next("this is not interesting") 
	b2.Mine(1)
	if b2.Proof != 40{
		t.Error("Block 2 Proof is not same as instructor proof")
	}
	if hex.EncodeToString(b2.Hash) != "ba2f9bf0f9ec629db726f1a5fe7312eb76270459e3f5bfdc4e213df9e47cd380"{
		t.Error("Block 2 Hash is not the same as the instructor hash")
	}
}


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
func TestMiningSpeed(t *testing.T){

	marginErr := 1.15

	//Run sequential mine for benchmark
	b0s := Initial(20)
	seqStart := time.Now()
	b0s.MineSequential()
	seqEnd := time.Now()
	targetTime := float64(seqEnd.Sub(seqStart))

	//Run concurrent miner with 1 worker (should be close to sequential)
	b0c := Initial(20)	
	concStart := time.Now()
	b0c.Mine(1)
	concEnd := time.Now()
	timeTaken := float64(concEnd.Sub(concStart))

	
	if timeTaken > targetTime * marginErr {
		t.Error("1 Worker concurrent mining is taking too long")
	}

	//Run concurrent miner with 4 workers, should be faster than sequential
	b1c := b0c.Next("Some Hash")
	concStart = time.Now()
	b1c.Mine(4)
	concEnd = time.Now()
	timeTaken = float64(concEnd.Sub(concStart))

	if timeTaken > targetTime { 
		t.Error ("Concurrent mining is slower than sequential")
	}

}


// TODO: some useful tests of Blocks
