package main

import (
	"github.com/mike-pittelko/sar/slab"
)

func main() {
	// Create some sample data

	var slab1 = slab.NewSlabWithSize(250000)
	slab1.RandomFill()

	slab1.Test()

}
