package main

import (
	"fmt"

	"time"

	"learngo/retriever/mock"
	"learngo/retriever/real"
)

/* 接口 */
type RetrieverItf interface {
	Get(url string) string
}
type PosterItf interface {
	Post(url string, form map[string]string) string
}
type RetrieverPoster interface {
	RetrieverItf
	PosterItf
}

const url = "http://www.imooc.com"

func download(r RetrieverItf) string {
	return r.Get(url)
}

func post(poster PosterItf) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

func session(s RetrieverPoster) string {
	s.Post(url,
		map[string]string{
			"contents": "another faked imooc.com",
		})
	return s.Get(url)
}

func main() {
	var itf RetrieverItf

	var mockRetriever = mock.Retriever{Contents: "this is a fake imooc.com"}
	itf = mockRetriever
	inspect(itf)

	itf = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(itf)

	// Type assertion
	if mockRetriever, ok := itf.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("itf is not a mock retriever")
	}

	fmt.Println("Try a session with mockRetriever")
	fmt.Println(session(mockRetriever)) // mockRetriever继承了两个接口，因此可以传进去
}

func inspect(itf RetrieverItf) {
	fmt.Println("Inspecting", itf)
	fmt.Printf(" > Type:%T Value:%v\n", itf, itf)
	fmt.Print(" > Type switch: ")

	switch v := itf.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
