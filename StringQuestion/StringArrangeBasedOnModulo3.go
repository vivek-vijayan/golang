package main

import (
	"fmt"
	"sort"
)

func main() {
	type dataStructure struct {
		text string
		size int
	}
	var dataString = []string{"Welcome", "Game", "Work", "Holiday", "Festival"}
	sort.Strings(dataString)
	var tempList = []dataStructure{}
	for _, each := range dataString {
		var temp dataStructure
		temp.text = each
		temp.size = len(each) % 3
		tempList = append(tempList, temp)

	}
	sort.Slice(tempList, func(p, q int) bool {
		return tempList[p].size < tempList[q].size
	})

	fmt.Println(tempList)
}
