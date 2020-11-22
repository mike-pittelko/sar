package main

//TODO: Make slab a container for the slab, and the shards, figure our how "new" should work on a struct
//TODO: Existing slab definition is more like a shard than a slab. shard needs, header, sum, etc.

func main() {
	// Create some sample data

	var slab1 = NewSlabWithSize(250000)
	slab1.RandomFill()

	/*
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

	*/

}
