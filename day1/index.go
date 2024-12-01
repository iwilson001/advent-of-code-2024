package main

import (
	"fmt"
	"math"
	"os"
	"slices"
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
	slices.Sort(distances1)
	slices.Sort(distances2)
	total := 0
	for i := 0; i < len(distances1); i++ {
		total += int(math.Abs(float64(distances1[i] - distances2[i])))
	}
	fmt.Print(total)
}
