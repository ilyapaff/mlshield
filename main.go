package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const constValues = 10000.0

func main() {
	args := os.Args
	value := constValues
	for i := 1; i < len(args); i++ {
		percent, err := strconv.Atoi(args[i])
		if err != nil {
			log.Fatalln(errors.New("В качестве параметров, принимает целочисленные значения, а было введено: " + args[i]))
		}
		value = PercentageChange(value, percent)
	}
	result := (value * 100) / constValues
	println(fmt.Sprintf("%.2f", 100-result) + "%")
}

func PercentageChange(value float64, percent int) float64 {
	return (value * (100 - float64(percent))) / 100
}
