package goutil

import (
	"fmt"
	"os"
	"json"
)

func bubbleSort[T any](arr *[]T, compare func(c1 T, c2 T) bool) {
	n := len(*arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if compare((*arr)[j], (*arr)[j+1]) {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

func IndexOf[T any, V any](arr []T, val V, compareFunc func(c1 T, c2 V) bool) int {
	for i, v := range arr {
		if compareFunc(v, val) {
			return i
		}
	}
	return -1
}

func prettyPrintStruct[T any](val *[]T) {
	for _, v := range *val {
		s, _ := json.Marshal(v)
		fmt.Println(string(s))
		fmt.Println()
	}
}

func loadJson[T any](FileName string, inp T) error {
	content, err := os.ReadFile(FileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &inp)
	if err != nil {
		return err
	}

	return nil
}