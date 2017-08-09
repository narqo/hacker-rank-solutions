package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
)

func main() {
	rawn, _ := ioutil.ReadAll(os.Stdin)
	n, _ := strconv.ParseInt(string(rawn), 10, 64)

	if n == 0 {
		fmt.Println(1)
		return
	}

	ret := big.NewInt(n)
	for n > 1 {
		n--
		ret = ret.Mul(ret, big.NewInt(n))
	}
	fmt.Println(ret)
}
