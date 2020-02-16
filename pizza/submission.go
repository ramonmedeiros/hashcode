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
    buffer := bufio.NewReader(fileDescriptor)

    // split string and get max number 
    line, _, _ := buffer.ReadLine()
    maxNumber = strings.Split(string(line), " ")

    // number of different pizzas
    line2, _, _ := buffer.ReadLine()
    sizePizzas = strings.Split(string(line2), " ")

    // close file
    fileDescriptor.Close()

    return maxNumber, sizePizzas
}

func main() {
    total, sizesPizza:= parseDataset(EXAMPLE)

    // total slices and types
    totalSlices, _ := strconv.Atoi(total[0])
    types, _ := strconv.Atoi(total[1])

    // convert string list to int
    var sizes = make([]int,types)

    for index, element := range sizesPizza {
        sizes[index], _ = strconv.Atoi(element)
    }
    fmt.Println("Max slices", totalSlices)
    fmt.Println("Sizes", sizes)
}

