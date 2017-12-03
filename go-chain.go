package main

import (
	"fmt"
	"bytes"
	"crypto/sha256"
	"time"
)

// will be used in the future 
// keep it simple for now 
type transaction struct {
	sender		string
	reciever	string
	amount		float32
}

type block struct {
	index			uint
	timestamp		uint32
	data			string
	previous_hash	[]byte
	hash 			[]byte
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

func (bl *block) get_hash() []byte {

	// Combine the headers 'data', 'previous_hash', and 'timestamp' and hash it	
	var buffer bytes.Buffer
	buffer.WriteString(bl.data)
	buffer.WriteString(string(bl.previous_hash))
	buffer.WriteString(fmt.Sprint(bl.timestamp))
	
	// create hash and set it to the current block
	h := sha256.New()
	h.Write(buffer.Bytes())
	return h.Sum(nil)
}

func create_block(data string, previous_block block) block {
	
	new_block := block {
		index: previous_block.index + 1, 
		timestamp: uint32(time.Now().Unix()), 
		data: data,
		previous_hash: previous_block.hash,
	}
	new_block.hash = new_block.get_hash()

	return new_block
}

func main() {
	bc := blockchain{}
	bc.add_block("Hello")
	bc.add_block("World")
	fmt.Println(bc.blocks);
}
