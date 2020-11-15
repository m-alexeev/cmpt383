package blockchain

import(
	"encoding/hex"
)


type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

// Create new initial (generation 0) block, not setting the .Hash value.
func Initial(difficulty uint8) Block {
	first := new(Block) 
	
}


// Create new block to follow this block, with provided data, not setting the .Hash value.
func (prev_block Block) Next(data string) Block {
	// * Create new Block 
	newBlock := Block(prev_block.PrevHash, prev_block.Generation + 1, prev_block.Difficulty, data) 
	return newBlock 
}

// String that we hash for this block.
func (blk Block) hashString() string {
	return blk.hashStringProof(blk.Proof)
}

// String that we hash for this block, if we had blk.Proof == proof.
func (blk Block) hashStringProof(proof uint64) string {

}

// Calculate hash as if blk.Proof == proof.
// Separated from .CalcHash so we can test many proof values without
// modifying the Block.
func (blk Block) calcHashProof(proof uint64) []byte {
	// TODO
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	return blk.calcHashProof(blk.Proof)
}

// Would this hash end in enough null bits, if blk.Proof == proof?
func (blk Block) validHashProof(proof uint64) bool {
	// TODO
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	return blk.validHashProof(blk.Proof)
}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
