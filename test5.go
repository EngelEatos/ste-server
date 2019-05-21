package main

import (
	"errors"
	"fmt"
	"reflect"
)

func Difference(a, b []interface{}) (diff []interface{}) {
	m := make(map[interface{}]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}

func ToInterfaceSlice(a interface{}) ([]interface{}, error) {
	k := reflect.ValueOf(a)
	if k.Kind() != reflect.Slice {
		return nil, errors.New("ToInterfaceSlice given a non-slice parameter")
	}
	var interfaceSlice []interface{} = make([]interface{}, len(k))
	for i, d := range k.Interface() {
		interfaceSlice[i] = d
	}
	return interfaceSlice, nil
}
func main() {
	fmt.Println(Difference([]int{5, 4, 3, 2, 1}, []int{6, 2, 1}))
	// 						 0000 0000 0000 0000 0001
}
