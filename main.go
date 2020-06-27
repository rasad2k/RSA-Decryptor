package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Usage: ./script <c number> <d number> <n number>")
		os.Exit(1)
	}

	c := new(big.Int)
	c.SetString(os.Args[1], 16)
	d := new(big.Int)
	d.SetString(os.Args[2], 16)
	n := new(big.Int)
	n.SetString(os.Args[3], 16)

	res := new(big.Int)
	res.Exp(c, d, n)
	hexa := toHex(res)
	ans, err := hex.DecodeString(hexa)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}
	fmt.Println(string(ans))
}

func toHex(n *big.Int) string {
	return fmt.Sprintf("%x", n)
}
