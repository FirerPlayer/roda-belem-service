package blooms

import (
	"errors"
	"fmt"
	"os"

	"github.com/bits-and-blooms/bloom/v3"
)

type BloomFilter struct {
	Filter bloom.BloomFilter
	Length uint
}

// func NewBloomFilter(length uint) *BloomFilter {
// 	return &BloomFilter{
// 		Filter: *bloom.NewWithEstimates(length, 0.01),
// 		Length: length,
// 	}
// }

//	func NewBloomFilterWithFilter(filter bloom.BloomFilter) *BloomFilter {
//		return &BloomFilter{
//			Filter: filter,
//			Length: filter.Cap(),
//		}
//	}
func NewBloomFilter(length uint) *BloomFilter {
	fileInfo, err := os.Stat("google_places_ids.bloom")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			os.Create("google_places_ids.bloom")
			return &BloomFilter{
				Filter: *bloom.NewWithEstimates(length, 0.01),
				Length: length,
			}
		}
		panic(err)
	}
	data, err := os.ReadFile(fileInfo.Name())
	if err != nil {
		panic(err)
	}
	m, k := bloom.EstimateParameters(uint(len(data)), 0.01)
	data64 := make([]uint64, len(data))
	for i := 0; i < len(data); i++ {
		data64[i] = uint64(data[i])
	}
	filter := bloom.FromWithM(data64, m, k)
	return &BloomFilter{
		Filter: *filter,
		Length: k,
	}
}

// Add adds a string key to the BloomFilter.
//
// key: the string to add to the BloomFilter.
func (b *BloomFilter) Add(key string) {
	b.Filter.Add([]byte(key))
}

func (b *BloomFilter) AddIfNotContains(key string) {
	if b.NotContains(key) {
		b.Add(key)
	}
}

// NotContains checks if the BloomFilter does not contain the given key.
//
// key: string representation of the key to check.
// bool: true if key is not contained in the BloomFilter, false (1% rate of false positive) otherwise.
func (b *BloomFilter) NotContains(key string) bool {
	return !b.Filter.Test([]byte(key))
}

func (b *BloomFilter) SaveBloomFilter() error {
	p, err1 := os.Getwd()
	if err1 != nil {
		return err1
	}
	fmt.Println(p)

	filterBinary, err := b.Filter.MarshalJSON()
	if err != nil {
		return errors.New("Failed to save bloom filter: " + err.Error())
	}
	fileInfo, err := os.Stat("google_places_ids.bloom")
	if err == nil {
		if os.IsNotExist(err) {
			return errors.New("File .bloom does not exist: " + err.Error())
		}
		return errors.New("Failed to save bloom filter: " + err.Error())
	}

	os.Remove(fileInfo.Name())
	err = os.WriteFile(fileInfo.Name(), filterBinary, 0666)
	if err != nil {
		return errors.New("Failed to save bloom filter: " + err.Error())
	}
	return nil
}
