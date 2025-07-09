package dynamicarray

import "errors"

type DynamicArray struct {
	data     []int
	capacity int
	length   int
}

func NewDynamicArray() *DynamicArray {
	initCap := 4
	return &DynamicArray{
		data:     make([]int, initCap),
		capacity: initCap,
		length:   0,
	}
}

func (d *DynamicArray) resize() {
	newCap := d.capacity * 2
	newData := make([]int, newCap)
	copy(newData, d.data)
	d.data = newData
	d.capacity = newCap
}

func (d *DynamicArray) Append(x int) {
	if d.length == d.capacity {
		d.resize()
	}
	d.data[d.length] = x
	d.length++
}

func (d *DynamicArray) Get(i int) (int, error) {
	if i < 0 || i >= d.length {
		return 0, errors.New("index out of bounds")
	}
	return d.data[i], nil
}

func (d *DynamicArray) Set(i, x int) error {
	if i < 0 || i >= d.length {
		return errors.New("index out of bounds")
	}
	d.data[i] = x
	return nil
}

func (d *DynamicArray) Size() int {
	return d.length
}

func (d *DynamicArray) PopBack() error {
	if d.length == 0 {
		return errors.New("array is empty")
	}
	d.length--
	return nil
}