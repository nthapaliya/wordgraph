package cdawg

import (
	"encoding/json"
	"io/ioutil"
)

func EncodeToBinary(md MDawg) []uint32 {
	outputbuffer := []uint32{}

	for i := range md {
		for j := range md[i] {
			outputbuffer = append(outputbuffer, uint32(md[i][j]))
		}
		if len(md[i]) == 0 {
			outputbuffer = append(outputbuffer, eolBitmask)
		}
	}
	return outputbuffer
}

func DecodeFromBinary(input []uint32) MDawg {
	md := MDawg{}
	row := []int{}

	for _, val := range input {
		eol := val&eolBitmask != 0
		row = append(row, int(val))
		if eol {
			md = append(md, row)
			row = []int{}
		}
	}
	return md
}

func WriteToFile(filename string, md MDawg) error {
	b, err := json.Marshal(md)

	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadFromFile(filename string) (MDawg, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	md := &MDawg{}
	err = json.Unmarshal(b, md)
	if err != nil {
		return nil, err
	}
	return *md, nil
}
