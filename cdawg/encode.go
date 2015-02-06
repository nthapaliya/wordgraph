package cdawg

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func EncodeToBinary(md MDawg) ([]byte, error) {
	out := []uint32{}

	for i := range md {
		for j := range md[i] {
			out = append(out, uint32(md[i][j]))
		}
	}

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, out)
	if err != nil {
		return nil, fmt.Errorf("binary.Write failed: %s", err)
	}
	return buf.Bytes(), nil
}

func DecodeFromBinary(input []byte) (MDawg, error) {
	buf := bytes.NewBuffer(input)
	flatmd := make([]uint32, len(input)/4)
	err := binary.Read(buf, binary.LittleEndian, flatmd)
	if err != nil {
		return nil, fmt.Errorf("binary.Read failed: %s", err)
	}

	md := MDawg{}
	row := []int{}

	for _, val := range flatmd {
		eol := val&eolBitmask != 0
		row = append(row, int(val))
		if eol {
			md = append(md, row)
			row = []int{}
		}
	}
	return md, nil
}

func MarshalJSON(filename string, md MDawg) error {
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

func UnmarshalJSON(filename string) (MDawg, error) {
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
