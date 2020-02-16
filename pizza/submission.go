package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

const EXAMPLE = "a_example.in"
const SMALL = "b_small.in"

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func parseDataset(filePath string) (maxNumber []string, numberPizzas []string) {

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
    numberPizzas = strings.Split(string(line2), " ")

    // close file
    fileDescriptor.Close()

    return maxNumber, numberPizzas
}

func main() {
    a, b := parseDataset(EXAMPLE)
    fmt.Println(a)
    fmt.Println(b)

}

