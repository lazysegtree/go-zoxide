# go-zoxide

<div align="center">

[![GoDoc](https://godoc.org/github.com/lazysegtree/go-zoxide?status.svg)](https://godoc.org/github.com/lazysegtree/go-zoxide) 
[![Go Report Card](https://goreportcard.com/badge/github.com/lazysegtree/go-zoxide)](https://goreportcard.com/report/github.com/lazysegtree/go-zoxide)

</div>


Golang wrapper for zoxide - a smarter cd command. Allowing easy integration of zoxide in your golang project.

## Requirements

[zoxide](https://github.com/ajeetdsouza/zoxide) needs to be installed.

By default, `go-zoxide` will look for `zoxide` binary in `$PATH`. 
Feature to specify another location will be added soon.

## Usage

```go
	zClient, err := zoxide.New()
	if err != nil {
		fmt.Printf("Error while initializing zoxide : %v\n", err)
		return
	}
	res, err := zClient.QueryAll(os.Args[1:]...)
	if err != nil {
		fmt.Printf("Error while fetching zoxide results : %v\n", err)
		return
	}

	for _, r := range res {
		fmt.Printf("Score : '%f', Path : '%v'\n", r.Score, r.Path)
	}

```
See `examples/` for more examples

## Motivation

- `DRY Principle`: Every Go project shouldn't reimplement zoxide command execution, parsing, and error handling.
- `Existing solutions are inadequate`: The only existing solution I could find - [sesh/zoxide](https://pkg.go.dev/github.com/joshmedeski/sesh/zoxide) has major limitations, is project-specific and comes with extra dependencies. It has limited API, no stability guarantees (internal package), and no documentation for external use
  - No standalone Go library provides clean API, comprehensive functionality, proper docs, and zero deps.

# Projects using this library

Note - Feel free to submit a PR or create an issue to have your project mentioned here.

- [![GitHub Repo stars](https://img.shields.io/github/stars/yorukot/superfile?style=plastic&label=superfile)](https://github.com/yorukot/superfile) : A modern and fancy terminal file manager. Feature loaded. Focussed on UI and usability.