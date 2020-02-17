package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

const EXAMPLE = "a_example.in"
const SMALL = "b_small.in"

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
    fmt.Println(numbers)
    return numbers[0], numbers[1], numbers[2:]
}

func main() {
    totalSlices, types, sizesPizza:= parseDataset(os.Args[1])

    // convert string list to int
    var sizes []int

    sum := 0
    index := 0
    item := 0
    for index = types -1; index >= 0; index-- {
        item = sizesPizza[index]
	if (item <= (totalSlices - sum) ) {
            sum = sum + item
	    sizes = append(sizes, index)
        }
    }
    fmt.Println("Sizes readed", len(sizesPizza))
    fmt.Println("Sizes length", types)
    fmt.Println("Max slices", totalSlices)
    fmt.Println("Sum", sum)
}

