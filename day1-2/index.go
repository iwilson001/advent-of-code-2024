package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	// fmt.Print(string(data))
	strData := string(data)
	firstSplit := strings.Split(strData, "\n")
	var distances1 []int
	var distances2 []int
	for _, el := range firstSplit {
		tmpSecondSplit := strings.Split(el, "   ")
		firstNum := tmpSecondSplit[0]
		secondNum := tmpSecondSplit[1]
		if i, err := strconv.Atoi(secondNum); err == nil {
			distances2 = append(distances2, i)
		}
		if i, err := strconv.Atoi(firstNum); err == nil {
			distances1 = append(distances1, i)
		}
	}
	total := 0
	countMap := make(map[int]int)
	for _, el := range distances1 {
		mapVal, mapExists := countMap[el]
		if mapExists {
			total += mapVal * el
		} else {
			countMap[el] = 0
			for _, el2 := range distances2 {
				if el == el2 {
					countMap[el] += 1
				}
			}
			total += el * countMap[el]
		}
	}
	fmt.Println(countMap)
	fmt.Println(total)
}
