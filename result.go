package zoxide

import (
	"strconv"
	"strings"
)

// Result represents a single result from a zoxide query.
// It contains the matched directory path and its score (frequency/recency).
type Result struct {
	Score float64 // The score assigned by zoxide (higher = more relevant)
	Path  string  // The full path to the directory
}

func parseResults(stdout string, withScore bool) ([]Result, error) {
	res := []Result{}

	// TODO: Is this the most efficient way of doing this ?
	for line := range strings.Lines(stdout) {
		line = strings.TrimSpace(line)
		// TODO: Move this to utility function for reusability
		var score float64
		var path string
		if withScore {
			score, path = splitScoreAndPath(line)
		} else {
			score, path = defaultScore, line
		}
		res = append(res, Result{Score: score, Path: path})
	}
	return res, nil
}

func splitScoreAndPath(line string) (float64, string) {
	tokens := strings.SplitN(line, scoreSeperator, 2)
	// TODO: Length should be 2

	// TODO : Handle error
	score, _ := strconv.ParseFloat(tokens[0], 64)

	return score, strings.TrimSpace(tokens[1])
}
