package dynamicarray

import "errors"

type DynamicArrayExtras struct {
	DynamicArray
}

func NewDynamicArrayExtras() *DynamicArrayExtras {
	return &DynamicArrayExtras{
		DynamicArray: *NewDynamicArray(),
	}
}

func (d *DynamicArrayExtras) Pop(i int) (int, error) {
	if i < 0 || i >= d.length {
		return 0, errors.New("index out of bounds")
	}
	val := d.data[i]
	for j := i; j < d.length-1; j++ {
		d.data[j] = d.data[j+1]
	}
	d.length--
	return val, nil
}

func (d *DynamicArrayExtras) Contains(x int) bool {
	for i := 0; i < d.length; i++ {
		if d.data[i] == x {
			return true
		}
	}
	return false
}

func (d *DynamicArrayExtras) Insert(i int, x int) error {
	if i < 0 || i > d.length {
		return errors.New("index out of bounds")
	}
	if d.length == d.capacity {
		d.resize()
	}
	for j := d.length; j > i; j-- {
		d.data[j] = d.data[j-1]
	}
	d.data[i] = x
	d.length++
	return nil
}

func (d *DynamicArrayExtras) Remove(x int) int {
	for i := 0; i < d.length; i++ {
		if d.data[i] == x {
			for j := i; j < d.length-1; j++ {
				d.data[j] = d.data[j+1]
			}
			d.length--
			return i
		}
	}
	return -1
}