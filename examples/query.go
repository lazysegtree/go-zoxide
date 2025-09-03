//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/lazysegtree/go-zoxide"
)

func main() {
	zClient, err := zoxide.New()
	if err != nil {
		fmt.Printf("Error while initializing zoxide : %v\n", err)
		return
	}
	res, err := zClient.Query(os.Args[1:]...)
	if err != nil {
		fmt.Printf("Error while fetching zoxide results : %v\n", err)
		return
	}
	fmt.Printf("Path : '%v'\n", res)

}
