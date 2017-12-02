package main

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
)

type transaction struct {
	sender		string
	reciever	string
	amount		float32
}

type block struct {

	index		uint
	timestamp	uint32
	transactions	transaction
	proof_of_work	string
	previous_hash	string
}

func main() {
	
	// testing creation of transaction and block
	trans := transaction{sender: "abc", reciever: "xyz", amount: 5.0}
	blk := block{index: 1, timestamp: 232, transactions: trans, previous_hash: "hi"}
	
	fmt.Println(blk)
	fmt.Println(trans)

	// testing creation of sha256
	s := "Welcome to the blockchain"
    	h := sha256.New()
    	h.Write([]byte(s))
    	sha1_hash := hex.EncodeToString(h.Sum(nil))
	
	// Hello :)
	fmt.Println("Hello, world")
	fmt.Println(sha1_hash)

}
