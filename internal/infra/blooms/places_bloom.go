package blooms

import "github.com/bits-and-blooms/bloom/v3"

type BloomFilter struct {
	Filter bloom.BloomFilter
	Length uint
}

func NewBloomFilter(length uint) *BloomFilter {
	return &BloomFilter{
		Filter: *bloom.NewWithEstimates(length, 0.01),
		Length: length,
	}
}

// Add adds a string key to the BloomFilter.
//
// key: the string to add to the BloomFilter.
func (b *BloomFilter) Add(key string) {
	b.Filter.Add([]byte(key))
}

// NotContains checks if the BloomFilter does not contain the given key.
//
// key: string representation of the key to check.
// bool: true if key is not contained in the BloomFilter, false (1% rate of false positive) otherwise.
func (b *BloomFilter) NotContains(key string) bool {
	return !b.Filter.Test([]byte(key))
}
