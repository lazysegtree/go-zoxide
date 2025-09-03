//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lazysegtree/go-zoxide"
)

func main() {
	temp_dir := filepath.Join(os.TempDir(), "example_dir")

	zClient, err := zoxide.New(zoxide.WithDataDir(temp_dir))
	if err != nil {
		fmt.Printf("Error while initializing zoxide : %v\n", err)
		os.Exit(1)
	}
	err = zClient.Add(temp_dir)
	if err != nil {
		fmt.Printf("Error while adding to zoxide : %v\n", err)
		os.Exit(2)
	}
	res, err := zClient.QueryAll(os.Args[1:]...)
	if err != nil {
		fmt.Printf("Error while fetching zoxide results : %v\n", err)
		os.Exit(3)
	}

	for _, r := range res {
		fmt.Printf("Score : '%f', Path : '%v'\n", r.Score, r.Path)
	}
}
