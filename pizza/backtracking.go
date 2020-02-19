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

// store best results
var best stackType
var bestSum int = 0

func (s *stackType) Push(v int) {
	*s = append(*s, v)
}

func (s *stackType) Pop() int {
	// FIXME: What do we do if the stack is empty, though?
	l := len(*s)
	ret := (*s)[l-1]
	*s = (*s)[:l-1]
	return ret
}

func check(e error) {
	if e != nil {
		fmt.Println("Error while processing file: ", e)
		os.Exit(1)
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

func Solve(begin int, array stackType, numbers []int, length int, target int, sum int) stackType {
	var i int
	if sum > bestSum {
		best = array
		bestSum = sum
	}

	if target == sum {
		return array
	}
	for i = begin; i < length; i++ {
		if sum+numbers[i] <= target {
			array.Push(numbers[i])
			Solve(i+1, array, numbers, length, target, sum+numbers[i])
			array.Pop()
		}
	}
	return best
}

func generateSubsets(numbers []int, length int, target int) stackType {
	var array stackType
	var sum int = 0
	stack := Solve(0, array, numbers, length, target, sum)
	if len(stack) == 0 {
		fmt.Println("No Solution")
	}
	fmt.Println("Sum:", bestSum)
	fmt.Println(stack)
	return stack
}

func main() {
	// error if file is not provided
	if len(os.Args) == 0 {
		fmt.Println("Specify the input file as argument")
		os.Exit(1)
	}
	totalSlices, length, sizesPizza := parseDataset(os.Args[1])
	fmt.Println("Max Slices:", totalSlices)
	sizes := generateSubsets(sizesPizza, length, totalSlices)
	saveAnswer(sizes, os.Args[1])
}
