package structs

import (
	"fmt"
	"strconv"
)

type Configuration struct {
	Arr    []bool
	Weight int
	Value  int

	Values  *[]int
	Weights *[]int
}

func (this *Configuration) ToString() string {
	str := strconv.Itoa(this.Value) + " " + strconv.Itoa(this.Weight)

	for _, elem := range this.Arr {
		if elem {
			str += " 1"
		} else {
			str += " 0"
		}
	}

	return str
}

func (this *Configuration) Print() {

	fmt.Println(this.ToString())
}

func (configuration *Configuration) Len() int {
	return len(configuration.Arr)
}

func NewConf(size int, values *[]int, weights *[]int) *Configuration {
	return &Configuration{make([]bool, size), 0, 0, values, weights}
}

func (this *Configuration) AddElement(index int) {
	if this.Arr[index] {
		return
	}

	this.Arr[index] = true

	this.Value += (*this.Values)[index]
	this.Weight += (*this.Weights)[index]
}

func (this *Configuration) Clone() *Configuration {
	res := NewConf(this.Len(), this.Values, this.Weights)

	res.Copy(this)
	return res
}

func (this *Configuration) Copy(other *Configuration) *Configuration {
	copy(this.Arr, other.Arr)
	this.Value = other.Value
	this.Weight = other.Weight
	this.Values = other.Values
	this.Weights = other.Weights
	return this
}
