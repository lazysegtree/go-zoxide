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
		os.Exit(1)
	}
	res, err := zClient.QueryAll(os.Args[1:]...)
	if err != nil {
		fmt.Printf("Error while fetching zoxide results : %v\n", err)
		os.Exit(2)
	}

	for _, r := range res {
		fmt.Printf("Score : '%f', Path : '%v'\n", r.Score, r.Path)
	}
}

/* Usage
➜  examples git:(main) ✗ go run query_all.go
Score : '84.000000', Path : '/Users/nitin/Programming/superfile'
Score : '20.000000', Path : '/Users/nitin/Programming'
Score : '20.000000', Path : '/Users/nitin/Programming/zoxide-go/examples'
Score : '16.000000', Path : '/Users/nitin/Programming/zoxide-go'
Score : '16.000000', Path : '/Users/nitin/Programming/spf_forks'
Score : '14.000000', Path : '/Users/nitin/Programming/carto'
Score : '4.000000', Path : '/Users/nitin/Library/Application Support/superfile'
Score : '2.000000', Path : '/Users/nitin/Programming/shared'
Score : '2.000000', Path : '/Users/nitin/Programming/spf_forks/iZarrios'
Score : '1.200000', Path : '/Users/nitin/Programming/superfile/src'
Score : '1.000000', Path : '/Users/nitin/Programming/spf_forks/magic/testsuite'
Score : '1.000000', Path : '/Users/nitin/Programming/spf_forks/magic'
Score : '0.500000', Path : '/System/Volumes/xarts'
Score : '0.200000', Path : '/Users/nitin/Programming/superfile/testsuite'
Score : '0.200000', Path : '/Users/nitin/Programming/shared/Misc/linux/ShellSources'
➜  examples git:(main) ✗ go run query_all.go prog fork
Score : '16.000000', Path : '/Users/nitin/Programming/spf_forks'
➜  examples git:(main) ✗
*/
