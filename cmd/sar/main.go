package main

import (
	"fmt"

	"math/rand"

	"crypto/sha1"

	"github.com/google/uuid"
	"github.com/klauspost/reedsolomon"
)

type slabHeader struct {
	slabIdent string
	slabHash  [sha1.Size]byte
}

// SlabSize ... : Default size of a slab
const SlabSize = 250000 // Default size of a slab

type slab struct {
	header    slabHeader
	slabBytes []byte
}

//TODO: Make slab a container for the slab, and the shards, figure our how "new" should work on a struct

func main() {
	// Create some sample data

	var slab1 = new(slab)
	slab1.header.slabIdent = uuid.New().String()
	slab1.slabBytes = make([]byte, SlabSize)
	slab1.header.slabHash = sha1.Sum(slab1.slabBytes)
	//	slab1.header.slabHash.Sum(slab1.slabBytes)

	for i := range slab1.slabBytes {
		slab1.slabBytes[i] = byte(rand.Intn(255))
	}

	enc, _ := reedsolomon.New(7, 4)

	shards, _ := enc.Split(slab1.slabBytes)

	_ = enc.Encode(shards)

	ok, _ := enc.Verify(shards)
	if ok {
		fmt.Println("SAR: Verify ok")
	}

	shards[4], shards[6], shards[7], shards[3] = nil, nil, nil, nil
	fmt.Println("SAR: Destroyed shards 4,6,7,3 (7 shards, 4 destroyed)")
	_ = enc.Reconstruct(shards)

	ok, _ = enc.Verify(shards)
	if ok {
		fmt.Println("SAR: Reconstruct ok")
	} else {
		fmt.Println("SAR: Failed to reconstruct")
	}
}
