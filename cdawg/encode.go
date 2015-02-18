package cdawg

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// EncodeToBytes encodes the CDawg to bytes for compact storage as a binary file
// on your hard-disk
//
func EncodeToBytes(cd CDawg) ([]byte, error) {
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

// DecodeFromBytes takes in a byte sequence that is read from a saved binary file
// representing the CDawg
//
func DecodeFromBytes(input []byte) (CDawg, error) {
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

// MarshalJSON saves the CDawg structure as a simple JSON file for cross
// program portability
//
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

// UnmarshalJSON reads a JSON file that has been previously saved. Note this is only
// for the 2D array form CDawg.
//
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

// reference functions, not exported yet. Just for testing
//
func _WriteMDtofile() {
	cd, err := UnmarshalJSON("../files/cd.json")
	mm, err := cd.Minimize()
	if err != nil {
		return
	}
	out := make([]uint32, len(mm))
	for i, v := range mm {
		out[i] = uint32(v)
	}

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, out)
	if err != nil {
		return
	}

	err = ioutil.WriteFile("md.bin.tmp", buf.Bytes(), 0644)
	if err != nil {
		return
	}
}

func _ReadMDfromfile() {
	b, err := ioutil.ReadFile("md.bin.tmp")
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(b)

	out := make([]uint32, buf.Len()/4)

	err = binary.Read(buf, binary.LittleEndian, out)
	if err != nil {
		return
	}
	fmt.Println(out[:20])

}
