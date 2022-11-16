package main

import (
	"encoding/csv"
	"io"
	"strconv"
)

func newDictionary(r io.Reader) (map[string]int, error) {
	reader := csv.NewReader(r)
	dict := make(map[string]int)
	for {
		arr, err := reader.Read()
		if err == io.EOF {
			return dict, nil
		}
		if err != nil {
			return nil, err
		}
		ans, err := strconv.Atoi(arr[1])
		if err != nil {
			return dict, err
		}
		dict[arr[0]] = ans
	}
}
