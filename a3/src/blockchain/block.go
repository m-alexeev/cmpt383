package blockchain

import(
	"encoding/hex"
	"strconv"
	"bytes"
	"crypto/sha256"
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
	first := Block{}
	first.PrevHash = bytes.Repeat([]byte("\x00"), 32)
	first.Generation = 0
	first.Difficulty = difficulty; 
	first.Data = ""
	return first
}


// Create new block to follow this block, with provided data, not setting the .Hash value.
func (prev_block Block) Next(data string) Block {
	// * Create new Block 
	newBlock := Block{}
	newBlock.PrevHash = prev_block.Hash
	newBlock.Generation = prev_block.Generation + 1
	newBlock.Difficulty = prev_block.Difficulty
	newBlock.Data =  data
	return newBlock 
}

// String that we hash for this block.
func (blk Block) hashString() string {
	return blk.hashStringProof(blk.Proof)
}

// String that we hash for this block, if we had blk.Proof == proof.
func (blk Block) hashStringProof(proof uint64) string {
	// *Create the output string
	byteStr := []byte(blk.PrevHash)
	encodedHash := hex.EncodeToString(byteStr)
	retString := "" + encodedHash + ":" + 
					strconv.FormatUint(blk.Generation,10) + ":" + 
					strconv.Itoa(int(blk.Difficulty)) + ":" + 
					blk.Data + ":" + 
					strconv.FormatUint(proof, 10)
	
	return retString
}

// Calculate hash as if blk.Proof == proof.
// Separated from .CalcHash so we can test many proof values without
// modifying the Block.
func (blk Block) calcHashProof(proof uint64) []byte {
	// TODO
	hash32 := sha256.Sum256([]byte(blk.hashString()))
	hash := hash32[:]
	return hash
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	return blk.calcHashProof(blk.Proof)
}

// Would this hash end in enough null bits, if blk.Proof == proof?
func (blk Block) validHashProof(proof uint64) bool {
	// TODO
	validHash := true
	nBytes := blk.Difficulty / 8
	nBits := blk.Difficulty % 8
	hash := blk.calcHashProof(proof)
	if len(hash) == 0{
		return false 
	}
	for i := uint8(8); i < nBytes ; i ++{
		if hash[len(hash) - int(i)- 1] != '\x00'{
			validHash = false
			break
		}
	}
	if  hash[len(hash) - int(nBytes) - 1] % (1<< nBits) != 0{
		validHash = false
	}

	return validHash
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


