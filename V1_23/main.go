package main

import (
	"fmt"
	"iter"
	"maps"
	"net/http"
	"slices"
)

func main() {
	SliceRepeat()
	MapFormal()
	HttpParseCookie()
	HttpParseSetCookie()
}

func SliceRepeat() {
	arr1 := []int{1, 2, 3}
	arr2 := slices.Repeat(arr1, 4)
	fmt.Printf("slices repeat %v, 4 times: %v\n", arr1, arr2)
	slices.Reverse(arr2)
	fmt.Println("slices reverse:", arr2)
}

func MapFormal() {
	m := map[string]struct{}{"張三": {}, "小黃": {}, "老吳": {}}

	for _, key := range slices.Sorted(maps.Keys(m)) {
		fmt.Printf("%s\t", key)
	}
	fmt.Println()

	// iter.Seq
	s1 := maps.Keys(m)
	for key := range s1 {
		fmt.Printf("%s\t", key)
	}
	fmt.Println()

	// iter key or values
	next, stop := iter.Pull(s1)
	defer stop()
	for {
		key, valid := next()
		if valid {
			fmt.Println(key)
		} else {
			break
		}
	}

	// iter (key,values) pairs
	s2 := maps.All(m)
	next2, stop2 := iter.Pull2(s2)
	defer stop2()
	for {
		k, v, valid := next2()
		if valid {
			fmt.Println(k, v)
		} else {
			break
		}
	}
}

// client to server, Cookie: session_id=abc123; MaxAge=0; lang=en; lang=zh-CN
func HttpParseCookie() {
	line := "session_id=abc123; value=hello-world; lang=en; lang=zh-CN"
	cookies, _ := http.ParseCookie(line) // new

	for _, cookie := range cookies {
		fmt.Printf("%v=%v\n", cookie.Name, cookie.Value)
	}
}

// server to client, Set-Cookie: session_id=abc123; MaxAge=0; lang=en; lang=zh-CN Domain=.abc.com
func HttpParseSetCookie() {
	line := "session_id=abc123; MaxAge=0; lang=en; lang=zh-CN Domain=.abc.com"
	cookie, _ := http.ParseSetCookie(line) // new

	fmt.Println("Raw: ", cookie.Raw)
	fmt.Println("Name: ", cookie.Name) // cookie[0]
	fmt.Println("Value: ", cookie.Value)
	fmt.Println("Domain: ", cookie.Domain)
	fmt.Println("MaxAge: ", cookie.MaxAge)
}
