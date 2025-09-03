//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/lazysegtree/go-zoxide"
)

// This file will be ignored during normal builds

func main() {

	if len(os.Args) > 1 {
		fmt.Println("Error : Provide non zero arguments")
	}
	zClient, err := zoxide.New()
	if err != nil {
		fmt.Printf("Error while initializing zoxide : %v\n", err)
		return
	}
	res, err := zClient.Query(os.Args...)
	if err != nil {
		fmt.Printf("Error while fetching zoxide results : %v\n", err)
		return
	}

	for _, r := range res {
		fmt.Printf("Score : '%f', Path : '%v'\n", r.Score, r.Path)
	}
}
