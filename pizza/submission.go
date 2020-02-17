package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

const EXAMPLE = "a_example.in"
const SMALL = "b_small.in"

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func parseDataset(filePath string) (maxNumber []string, sizePizzas []string) {

    // open fd
    fileDescriptor, err := os.Open(filePath)
    check(err)

    // create buffer for read lines
    buffer := bufio.NewReaderSize(fileDescriptor, 1024 * 10)

    // split string and get max number 
    line, _, _ := buffer.ReadLine()
    maxNumber = strings.Split(string(line), " ")

    // number of different pizzas
    buffer.Reset()
    line2, _, _ := buffer.ReadLine()
    sizePizzas = strings.Split(string(line2), " ")

    // close file
    fileDescriptor.Close()

    return maxNumber, sizePizzas
}

func main() {
    total, sizesPizza:= parseDataset(os.Args[1])

    // total slices and types
    totalSlices, _ := strconv.Atoi(total[0])
    types, _ := strconv.Atoi(total[1])

    // convert string list to int
    var sizes []int

    sum := 0
    index := 0
    item := 0
    for index = types -1; index >= 0; index-- {
        item, _ = strconv.Atoi(sizesPizza[index])
	if (item <= (totalSlices - sum) ) {
            sum = sum + item
	    sizes = append(sizes, index)
        }
    }
    fmt.Println("Sizes readed", len(sizes))
    fmt.Println("Sizes length", types)
    fmt.Println("Max slices", totalSlices)
    fmt.Println("Sum", sum)
}

