package main

import "fmt"
import "math/rand"
import "github.com/klauspost/reedsolomon"

func main() {
	// Create some sample data
	var data = make([]byte, 250000)

	for i := range data {
		data[i] = byte(rand.Intn(8))
	}

	enc, _ := reedsolomon.New(7, 3)

	shards, _ := enc.Split(data)

	_ = enc.Encode(shards)

	ok, _ := enc.Verify(shards)
	if ok {
		fmt.Println("Verify ok")
	}

	shards[4], shards[6] = nil, nil

	_ = enc.Reconstruct(shards)

	ok, _ = enc.Verify(shards)
	if ok {
		fmt.Println("Reconstruct ok")
	}
}
