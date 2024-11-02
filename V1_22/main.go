package main

import (
	"fmt"
	"net/http"
	"slices"
	"sync"
)

func main() {
	LoopVar()
	LoopIntRange()
	SlicesConcat()
	HttpService()
}

func LoopVar() {
	wg := sync.WaitGroup{}
	values := []string{"a", "b", "c", "d"}
	wg.Add(len(values))

	for _, val := range values {
		// NOW: val will reallocate memory every time
		fmt.Printf("%v: %p\n", val, &val)

		go func() {
			defer wg.Done()
			fmt.Printf("goroutine-> %v: %p\n", val, &val)
		}()
	}
	wg.Wait()
}

func LoopIntRange() {
	// start from 0 ~ 5
	for num := range 6 {
		fmt.Println("loop int range:", num)
	}
}

func SlicesConcat() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	arr3 := []int{7, 8, 9}
	fmt.Println("slices concat:", slices.Concat(arr1, arr2, arr3))
}

func HttpService() {
	mux := http.NewServeMux()

	// restful參數
	mux.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			fmt.Fprintln(w, "非GET Method")
		} else {
			fmt.Println("received path value \"name\":", r.PathValue("name"))
			fmt.Fprintf(w, "Hello %s!!!\n", r.PathValue("name"))
		}
	})

	// 直接指定HTTP Method, 錯誤返回405 Method Not Allowed
	mux.HandleFunc("POST /gogogo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("received post request!")
		fmt.Fprintln(w, "gogogo!!!")
	})

	fmt.Println("http server start...")
	http.ListenAndServe("127.0.0.1:8080", mux)
}
