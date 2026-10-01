package main

import (
	"fmt"
	"os"

	"github.com/ibrahimamar07/devctl/cmd"
)

func main() {
	cmd.Execute(os.Args[1:])

	x := 10

	p := &x

	*p = 20

	fmt.Println(*p, x)
}
