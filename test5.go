package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

// func sliceNdice(chapters []int) [][]int {
// 	vCount := (len(chapters) + 100 - 1) / 100
// 	result := make([][]int, vCount)
// 	for i := 0; i < vCount; i++ {
// 		eIdx := (i + 1) * 100
// 		if eIdx > len(chapters) {
// 			eIdx = len(chapters)
// 		}
// 		fmt.Printf("idx: %d, sidx: %d, eidx: %d\n", i, i*100, eIdx)
// 		result[i] = chapters[i*100 : eIdx]
// 	}
// 	return result
// }
func sliceNdice(input interface{}, chunksize int) ([]interface{}, error) {
	s := reflect.ValueOf(input)
	if s.Kind() != reflect.Slice {
		return nil, errors.New("sliceNdice given a non-slice type")
	}
	length := s.Len()
	vCount := (length + chunksize - 1) / chunksize
	result := make([]interface{}, vCount)
	for i := 0; i < vCount; i++ {
		eIdx := (i + 1) * chunksize
		if eIdx > length {
			eIdx = length
		}
		r := s.Slice(i*chunksize, eIdx).Interface()
		result[i] = r
	}
	return result, nil
}

func main() {
	var k []int
	for i := 1; i <= 950; i++ {
		k = append(k, i)
	}
	r, err := sliceNdice(k, 100)
	if err != nil {
		log.Fatal(err)
	}
	f := reflect.ValueOf(r)
	fmt.Println(f.Type())
	o, ok := f.Interface().([][]int)
	if !ok {
		panic("failed")
	}
	for _, l := range o {
		fmt.Println(l)
	}

}
