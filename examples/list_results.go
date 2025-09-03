//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"zoxide"
)

// This file will be ignored during normal builds

func main() {

	if len(os.Args) > 1 {
		fmt.Println("Error : Provide non zero arguments")
	}
	zClient := zoxide.New()
	res, err := zClient.Query(os.Args...)
	if err != nil {
		fmt.Println("Error while fetching zoxide results : %v", err)
		return
	}

	for _, r := range res {
		fmt.Printf("Score : '%f', Path : '%v'\n", r.Score, r.Path)
	}
}
