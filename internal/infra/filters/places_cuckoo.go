package filters

import (
	"fmt"
	"os"

	cuckoo "github.com/panmari/cuckoofilter"
)

const originalFileName = "tmp/google_places_ids.cuckoo"

type CuckooFilter struct {
	Filter cuckoo.Filter
	Length uint
}

func NewCuckooFilter(length uint) (*CuckooFilter, error) {
	_, err := os.Stat(originalFileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("- creating file %s\n", originalFileName)
			filter := cuckoo.NewFilter(length)
			err = os.WriteFile(originalFileName, filter.Encode(), 0666)
			if err != nil {
				return nil, fmt.Errorf("error creating file %s: %v", originalFileName, err.Error())
			}
			fmt.Printf("- file %s created\n", originalFileName)
			return &CuckooFilter{
				Filter: *filter,
				Length: length,
			}, nil
		}
		return nil, fmt.Errorf("error opening file %s: %v", originalFileName, err.Error())
	}
	fmt.Printf("File %s exists. Looking for filter data...\n", originalFileName)
	fileBinary, err := os.ReadFile(originalFileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", originalFileName, err.Error())
	}
	filter, err := cuckoo.Decode(fileBinary)
	if err != nil {
		return nil, fmt.Errorf("error decoding file %s: %v", originalFileName, err.Error())
	}
	fmt.Println("- data successfully loaded")

	return &CuckooFilter{
		Filter: *filter,
		Length: length,
	}, nil
}

// Add adds a string key to the CuckooFilter.
//
// key: the string to add to the CuckooFilter.
func (c *CuckooFilter) Add(key string) {
	c.Filter.Insert([]byte(key))
}

func (c *CuckooFilter) AddIfNotContains(key string) {
	if c.NotContains(key) {
		c.Add(key)
	}
}

// NotContains checks if the CuckooFilter does not contain the given key.
//
// key: string representation of the key to check.
// bool: true if key is not contained in the CuckooFilter, false (0,1% rate of false positive) otherwise.
func (c *CuckooFilter) NotContains(key string) bool {
	return !c.Filter.Lookup([]byte(key))
}

func (c *CuckooFilter) SaveCuckooFilter() error {
	filterBinary := c.Filter.Encode()
	err := os.Remove(originalFileName)
	if !os.IsNotExist(err) && err != nil {
		return fmt.Errorf("error removing file: %v", err)
	}

	err = os.WriteFile(originalFileName, filterBinary, 0666)
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}
	fmt.Println("Saved cuckoo filter to: ", originalFileName)

	return nil
}
