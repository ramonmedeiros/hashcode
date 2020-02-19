package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// define stacktype
type stackType []int

// push to list
func (s *stackType) Push(v int) {
	*s = append(*s, v)
}

// pop from list
func (s *stackType) Pop() int {
	// FIXME: What do we do if the stack is empty, though?
	l := len(*s)
	ret := (*s)[l-1]
	*s = (*s)[:l-1]
	return ret
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseDataset(filePath string) (int, int, []int) {

	// open fd
	fileDescriptor, err := os.Open(filePath)
	check(err)
	defer fileDescriptor.Close()

	// create buffer for read words
	scanner := bufio.NewScanner(fileDescriptor)
	scanner.Split(bufio.ScanWords)

	// read file
	var numbers []int
	for scanner.Scan() {
		item, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, item)
	}
	return numbers[0], numbers[1], numbers[2:]
}

func saveAnswer(array []int, filename string) {
	// create file
	s := []string{filename, ".answer"}
	saveFile := strings.Join(s, "")
	f, err := os.Create(saveFile)
	check(err)
	defer f.Close()

	t := strconv.Itoa(len(array)) + "\n"
	f.WriteString(t)
	for x := 0; x < len(array); x++ {
		v := strconv.Itoa(array[x]) + " "
		f.WriteString(v)
	}
	f.WriteString("\n")
	fmt.Println("File saved at", saveFile)
}

func sumUntilLimit(array []int, size int, target int) (stackType, int) {
	var sizes, bestSizes stackType
	var index, sum, temp, bestSum int

	sum = 0
	bestSum = 0
	for (len(sizes) > 0 && sizes[0] != 0) || len(sizes) == 0 {

		// each new iteration, remove the biggest number
		size = size - 1

		for index = size; index >= 0; index-- {
			temp = sum + array[index]
			if temp <= target {
				sizes.Push(index)
				sum = temp
				if sum == target {
					return sizes, sum
				}
			}
		}

		// save better solutions until now
		if sum > bestSum {
			bestSum = sum

			// HIDDEN MAGIC SPOTTED: cannot copy with copy() or =
			for _, element := range sizes {
				bestSizes.Push(element)
			}
		}

		// remove the last number to try smaller ones
		if len(sizes) != 0 {
			last := sizes.Pop()
			sum = sum - array[last]
			size = last
		}

		// end of numbers to check
		if len(sizes) == 0 && size == 0 {
			break
		}
	}
	return bestSizes, bestSum
}

func main() {
	totalSlices, types, sizesPizza := parseDataset(os.Args[1])
	fmt.Println("Max Slices:", totalSlices)
	sizes, sum := sumUntilLimit(sizesPizza, types, totalSlices)
	fmt.Println("Sum:", sum)
	saveAnswer(sizes, os.Args[1])
}
