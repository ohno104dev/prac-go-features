package main

import (
	"context"
	"fmt"
	"log/slog"
	"slices"

	"golang.org/x/exp/maps"
)

func main() {
	BuildInFunc()
	Slog()
	SlicesFormal()
	MapsExp()
}

func BuildInFunc() {
	// min, max, clear
	fmt.Println("min:", min(1, 4, 7, 5, 9))
	fmt.Println("max:", max(1, 4, 7, 5, 9))
	list := []int{1, 3, 5, 7}
	m := map[string]string{"name": "小白"}
	clear(list)
	fmt.Printf("content: %v , len: %d , cap: %d\n", list, len(list), cap(list))
	clear(m)
	fmt.Printf("content: %v , len: %d\n", m, len(m))
}

func Slog() {
	slog.Log(context.Background(), slog.LevelInfo, "slog test", "姓名", "安安") //key value pairs
}

func SlicesFormal() {
	list := []int{12, 4, 6, 8, 100}
	fmt.Println("slice max:", slices.Max(list))
	fmt.Println("slice min:", slices.Min(list))
	fmt.Printf("slice: %v , 是否包含3: %t\n", list, slices.Contains(list, 3))
	list2 := slices.Clone(list)
	fmt.Printf("slice 深拷貝:\n  list:  %v, addr: %p\n  list2: %v, addr: %p\n", list, &list, list2, &list2)
	fmt.Println("slice equal(list1, list2):", slices.Equal(list, list2))
	slices.Sort(list)
	fmt.Println("slice sort:", list)
	index, exists := slices.BinarySearch(list, 6)
	fmt.Printf("slice search: 6是否存在: %t, index: %d\n", exists, index)
}

func MapsExp() {
	mp := map[string]string{"姓名": "小六", "性別": "男性", "班級": "3-4"}
	fmt.Printf("keys %s\n", maps.Keys(mp))
	fmt.Printf("Values %s\n", maps.Values(mp))
}
