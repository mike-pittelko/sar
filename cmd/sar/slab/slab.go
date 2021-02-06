package slab

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/google/uuid"
	"github.com/klauspost/reedsolomon"
)

// SlabHeader is the header for a slab, include the identity of the slab, the hash of the slab and pointers to the shards.
type slabHeader struct {
	slabIdent   string
	slabHash    [sha1.Size]byte
	shardHashes [][]byte
	shards      []Shard
	shardCount  uint64
	shardIdents [][]byte
}

// ShardHeader is the header for a single shard, a slice of a slab, that is to be managed with error coding
type shardHeader struct {
	shardIdent string
	shardHash  [sha1.Size]byte
}

// Slab is a file like object that contains the raw data, and the shards ready to be sent.
type Slab struct {
	header    slabHeader
	slabBytes []byte
}

// Shard is a segment of a slab
type Shard struct {
	header     shardHeader
	shardBytes []byte
}

// NewSlab is a factory function or empty slabs that are initialized with a UUID only.
func NewSlab() *Slab {
	slab1 := new(Slab)
	slab1.header.slabIdent = uuid.New().String()
	return slab1
}

// GetSlabBytes is a getter for the raw byte slice of the content of this slab
func (slab1 Slab) GetSlabBytes() []byte {
	return slab1.slabBytes
}

// NewSlabWithSize is a factory function for slabs that have a buffer preallocates with newSize uint64 bytes of zeros.
func NewSlabWithSize(newSize uint64) *Slab {
	slab1 := NewSlab()
	slab1.slabBytes = make([]byte, newSize)
	slab1.header.slabHash = sha1.Sum(slab1.slabBytes)
	return slab1
}

// NewSlabWithFile is a factory function for slabs that loads a file into the buffer
func NewSlabWithFile(file string) *Slab {
	slab1 := NewSlab()

	//	slab1.slabBytes = make([]byte, newSize)
	slab1.slabBytes, _ = ioutil.ReadFile(file)
	slab1.header.slabHash = sha1.Sum(slab1.slabBytes)
	return slab1
}

// RandomFill fills a slab (leaving shards untouched) to it's current size with random bytes. This is normally only useful in debugging scenarios
func (slab1 *Slab) RandomFill() *Slab {
	for i := range slab1.slabBytes {
		slab1.slabBytes[i] = byte(rand.Intn(255))
	}
	slab1.header.slabHash = sha1.Sum(slab1.slabBytes)
	return slab1
}

// SplitShards causes the shards to be generated from the data by using the configured number of endpoints
func (slab1 *Slab) SplitShards() *Slab {
	// use the enc.split here
	return slab1
}

// AddTarget adds a destination endpoint (at this time, cannot delete these) A better way is to create a
// "Target" that contains muultiple endpoints, and add that to the slab.
func (slab1 *Slab) AddTarget(endpointString string) *Slab {

	return slab1
}

//

/*https://golangcode.com/mocking-s3-upload/ */

// Test performs a simple test to demonstrate some reedsolomon stuff.  Go make a better test after RS has been incorporated into the slab properly
func (slab1 *Slab) Test() {

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

//todo: Build a "target" that contains multiple endpoints
/*

 */
