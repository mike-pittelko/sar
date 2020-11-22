package main

import (
	"crypto/sha1"
	"math/rand"

	"github.com/google/uuid"
)

// SlabHeader is the header for a slab, include the identity of the slab, the hash of the slab and pointers to the shards.
type SlabHeader struct {
	slabIdent string
	slabHash  [sha1.Size]byte
}

// ShardHeader is the header for a single shard, a slice of a slab, that is to be managed with error coding
type ShardHeader struct {
	shardIdent string
	shardHash  [sha1.Size]byte
}

// Slab is a file like object that contains the raw data, and the shards ready to be sent.
type Slab struct {
	header    SlabHeader
	slabBytes []byte
}

// Shard is a segment of a slab
type Shard struct {
	header     ShardHeader
	shardBytes []byte
}

// NewSlab is a factory function or empty slabs that are initialized with a UUID only.
func NewSlab() *Slab {
	slab1 := new(Slab)
	slab1.header.slabIdent = uuid.New().String()
	return slab1
}

// NewSlabWithSize is a factory function for slabs that have a buffer preallocates with newSize uint64 bytes of zeros.
func NewSlabWithSize(newSize uint64) *Slab {
	slab1 := NewSlab()
	slab1.slabBytes = make([]byte, newSize)
	slab1.header.slabHash = sha1.Sum(slab1.slabBytes)
	return slab1
}

// RandomFill fills a slab (leaving shards untouched) to it's current size with random bytes. This is normally only useful in debugging scenarios
func (x *Slab) RandomFill() *Slab {
	for i := range x.slabBytes {
		x.slabBytes[i] = byte(rand.Intn(255))
	}
	x.header.slabHash = sha1.Sum(x.slabBytes)
	return x
}

//TODO: Make slab a container for the slab, and the shards, figure our how "new" should work on a struct
//TODO: Existing slab definition is more like a shard than a slab. shard needs, header, sum, etc.
