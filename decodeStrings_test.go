package decodestrings

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"testing"
)

type Test struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

func TestDecodeString(tst *testing.T) {
	f, err := os.Open("./test.json")

	if err != nil {
		tst.Error(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	decoder := json.NewDecoder(reader)
	var tests map[string]Test
	for {
		err := decoder.Decode(&tests)

		if err == nil {
			for name, test := range tests {
				tst.Run(name, func(st *testing.T) {
					if DecodeString(test.Input) != test.Output {
						st.Errorf("Use case %v\n", test)
					}
				})
			}
		} else if err == io.EOF {
			break
		} else {
			tst.Error(err)
		}
	}

}
