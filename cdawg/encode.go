package cdawg

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func EncodeToBinary(cd CDawg) ([]byte, error) {
	out := []uint32{}

	for i := range cd {
		for j := range cd[i] {
			out = append(out, uint32(cd[i][j]))
		}
	}

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, out)
	if err != nil {
		return nil, fmt.Errorf("binary.Write failed: %s", err)
	}
	return buf.Bytes(), nil
}

func DecodeFromBinary(input []byte) (CDawg, error) {
	buf := bytes.NewBuffer(input)
	flatcd := make([]uint32, len(input)/4)
	err := binary.Read(buf, binary.LittleEndian, flatcd)
	if err != nil {
		return nil, fmt.Errorf("binary.Read failed: %s", err)
	}

	cd := CDawg{}
	row := []int{}

	for _, val := range flatcd {
		eol := val&eolBitmask != 0
		row = append(row, int(val))
		if eol {
			cd = append(cd, row)
			row = []int{}
		}
	}
	return cd, nil
}

func MarshalJSON(filename string, cd CDawg) error {
	b, err := json.Marshal(cd)

	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func UnmarshalJSON(filename string) (CDawg, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	cd := &CDawg{}
	err = json.Unmarshal(b, cd)
	if err != nil {
		return nil, err
	}
	return *cd, nil
}
