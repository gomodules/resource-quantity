package main

import (
	"encoding/json"
	"fmt"
	"gomodules.xyz/resource-quantity"
)

func main() {
	q := resource.MustParse("1.1Gi")
	if data, err := json.Marshal(q); err != nil {
		panic(err)
	} else {
		fmt.Println(string(data))
	}
}
