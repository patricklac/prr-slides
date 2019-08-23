// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START1 OMIT
func GoogleSearch(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

// STOP1 OMIT

// START2 OMIT
var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
	// ... other sources
)

type Result string // HL

type Search func(query string) Result // HL

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// STOP2 OMIT

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	start := time.Now()

	results := GoogleSearch("golang") // HL

	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed)
}
