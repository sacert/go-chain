package main

import (
	"fmt"
	"time"
	"crypto/sha256"
	"encoding/hex"
)

// sample of how a transaction would look
// this would replace the 'data' in block
type transaction struct {
	sender		string
	reciever	string
	amount		float32
}

type block struct {
	index			uint		// depth of current block
	timestamp		uint32		// timestamp in unix
	data			string		// can be anything - would be transaction
	previous_hash	string		// keep track of the last block 
	hash 			string		// hash of current block parameters, used sha256
	nonce			uint64		// used for proof of work
}

type blockchain struct {
	blocks []block
}

// check to see if the block is the genisis block otherwise add the block
// gensis block is the first block within the chain and will not be able
// to use the previous block for hashing or contain previous_hash
func (chain *blockchain) add_block(data string) {

	// check if the genesis block
	previous_block := block{}
	if (len(chain.blocks) != 0) {
		previous_block = chain.blocks[len(chain.blocks)-1]
	}
	new_block := create_block(data, previous_block)
	chain.blocks = append(chain.blocks, new_block)
}

func create_block(data string, previous_block block) block {
	
	new_block := block {
		index: previous_block.index + 1, 
		timestamp: uint32(time.Now().Unix()), 
		data: data,
		previous_hash: previous_block.hash,
		nonce: 0,
	}
	new_block.hash = new_block.get_hash()

	// Proof of Work
	// Find a hash that has the first 2 values equal to "00"
	// Should be a race by other nodes to find this node first
	// Exponential growth - makes it very hard to calculate depending
	// on the number of 0s (aka the difficulty level)
	for (new_block.hash[0:2] != "00") {
		new_block.nonce += 1
		new_block.hash = new_block.get_hash()
	}
	return new_block
}

func (bl *block) get_hash() string {
	
	// Combine the headers 'data', 'previous_hash', and 'timestamp' and hash it	
	s := bl.data + bl.previous_hash + fmt.Sprint(bl.timestamp) + string(bl.nonce)

	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	bc := blockchain{}
	bc.add_block("Hello")
	bc.add_block("World")
	
	fmt.Println("------");
	i := 0
	for (i < len(bc.blocks)) { 
		fmt.Println("Index: ", bc.blocks[i].index);
		fmt.Println("Timestamp: ", bc.blocks[i].timestamp);
		fmt.Println("Previous Hash: ", bc.blocks[i].previous_hash);
		fmt.Println("Hash: ", bc.blocks[i].hash);
		fmt.Println("Nonce: ", bc.blocks[i].nonce);
		fmt.Println("------");
		i++;
	}
}
