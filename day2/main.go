package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var safeReports int
var diff float64
var j int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var numbers [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Fields(line)
		var row []int
		for _, col := range cols {
			num, err := strconv.Atoi(col)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			row = append(row, num)
		}
		numbers = append(numbers, row)
	}
myLoop:
	for _, row := range numbers {
		if sort.IntsAreSorted(row) {
			for j, val := range row {
				if j+1 < len(row) {
					diff = float64(val - row[j+1])
					if math.Abs(diff) >= 1 && math.Abs(diff) <= 3 {
					} else {
						continue myLoop
					}
				}
			}
			safeReports = safeReports + 1
		}
	}
	fmt.Println("Result:", safeReports)
}
