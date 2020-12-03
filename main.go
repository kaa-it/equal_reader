package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
)

type EReader struct {
	total int
}

func (er *EReader) Read(b []byte) (int, error) {
	rand.Seed(99)

	if er.total > 2888 {
		return 0, io.EOF
	}

	maxCount := rand.Intn(1000) + 1

	if er.total + maxCount > 2888 {
		maxCount = 2888 - er.total
	}

	count := int(math.Min(float64(len(b)), float64(maxCount)))

	for i := 0; i < count; i++ {
		b[i] = 8
	}

	er.total = er.total + count

	if er.total == 2888 {
		return count, io.EOF
	}

	return count, nil
}

var r = &EReader{0}

func equalRead() {
	var buffer [1000]byte
	var count int
	var total int
	var err error

	for {
		count, err = r.Read(buffer[total:])
		if count == 0 {
			if err == io.EOF {
				if total > 0 {
					fmt.Printf("Read %d", total) // buffer[:total]
				}
				break
			}
		}

		if count == 1000 {
			fmt.Println("Read 1000") // buffer[:]
			count = 0
			continue
		}

		total = total + count

		if total == 1000 {
			fmt.Println("Read 1000") // buffer[:]
			total = 0
			continue
		}
	}
}

func main() {
	equalRead()
}