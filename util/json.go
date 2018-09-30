package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func Encode(v interface{}) []byte {
	b, err :=  json.Marshal(v)
	if err != nil {
		fmt.Printf("Encode err: %s\n", err)
		os.Exit(-1)
	}
	return b
}

func Decode(d []byte, v interface{}) {
	if err := json.Unmarshal(d, v); err != nil {
		fmt.Printf("Decode err: %s\n", err)
		os.Exit(-1)
	}
}
