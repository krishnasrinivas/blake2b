package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/minio/blake2b-simd"
)

func main() {
	input := os.Stdin
	if len(os.Args) == 2 {
		var err error
		input, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	h := blake2b.New512()
	_, _ = io.Copy(h, input)
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}
