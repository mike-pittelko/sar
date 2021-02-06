package slab

import (
	"fmt"
	"testing"

	"github.com/klauspost/reedsolomon"
)

//

/*https://golangcode.com/mocking-s3-upload/ */

// Test performs a simple test to demonstrate some reedsolomon stuff.  Go make a better test after RS has been incorporated into the slab properly
func Test(t *testing.T) {

	slab1 := NewSlabWithSize(2500000)
	slab1.RandomFill() // Fill with random data

	enc, _ := reedsolomon.New(7, 4)

	shards, _ := enc.Split(slab1.GetSlabBytes())

	_ = enc.Encode(shards)

	ok, _ := enc.Verify(shards)
	if ok {
		fmt.Println("SAR: Verify ok")
	} else {
		fmt.Println("SAR: Failed to Verify stage 1")
		t.Error("SAR: Failed to Verify stage 1 during Shard Test w/ RS")
	}

	shards[4], shards[6], shards[7], shards[3] = nil, nil, nil, nil
	fmt.Println("SAR: Destroyed shards 4,6,7,3 (7 shards, 4 destroyed)")
	_ = enc.Reconstruct(shards)

	ok, _ = enc.Verify(shards)
	if ok {
		fmt.Println("SAR: Reconstruct ok")
	} else {
		fmt.Println("SAR: Failed to verify reconstruction, stage 2")
		t.Error("SAR: Failed to verify reconstruction, stage 2, during Shard Test w/ RS")
	}

}

//todo: Build a "target" that contains multiple endpoints
/*

 */
