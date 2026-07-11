package main

import (
	"fmt"
	"strings"
)

/*A map maps keys to values.
The zero value of a map is nil. A nil map has no keys, nor can keys be added.
The make function returns a map of the given type, initialized and ready for use. */

type Vertexs struct {
	Lat, Long float64
}

// Map literals are like struct literals, but keys are required.
var m = map[string]Vertexs{
	"Bell Labs": Vertexs{
		40.68433, -74.39967,
	},
	"Google": Vertexs{
		37.42202, -122.08408,
	},
}

/*Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure. */
func WordCount(s string) map[string]int {
	wc := make(map[string]int, 0)
	strArr := strings.Fields(s)
	for i := range strArr {
		wc[strArr[i]] = wc[strArr[i]] + 1
	}
	return wc
}

func mapsDemo() {
	fmt.Println(m)
}
