package main

import (
	"fmt"
	"gomodules.xyz/resource-quantity"
)

func main() {
	q := resource.MustParse("1.1Gi")
	fmt.Println(q.String())
}
