package main

import "fmt"

type transaction struct {
	sender		string
	reciever	string
	amount		float32
}

type block struct {

	index		uint
	timestamp	uint32
	transactions	transaction
	previous_hash	string
}

func main() {
	trans := transaction{sender: "abc", reciever: "xyz", amount: 5.0}
	blk := block{index: 1, timestamp: 232, transactions: trans, previous_hash: "hi"}
	fmt.Println(blk)
	fmt.Println(trans)
	fmt.Printf("Hello, world")

}
