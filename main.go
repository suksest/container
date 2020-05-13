package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//Container contains balls
type Container struct {
	Capacity uint64
	Quantity uint64
}

func (c *Container) isFull() bool {
	if c.Quantity == c.Capacity {
		return true
	}
	return false
}

func (c *Container) isEmpty() bool {
	if c.Quantity == 0 {
		return true
	}

	return false
}

func (c *Container) fill() {
	if c.Quantity < c.Capacity {
		c.Quantity++
	}
}

func getRandomNumber(max, min int) int {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(max-min+1) + min)
}

func main() {
	n, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("Expected 2 parameters, got %v\n", len(os.Args))
		fmt.Println("Usage go run main.go [number_container] [number_capacity]")
	}
	cap, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		fmt.Printf("Expected 2 parameters, got %v\n", len(os.Args))
		fmt.Println("Usage go run main.go [number_container] [number_capacity]")
	}
	var arrContainer []Container
	for i := uint64(0); i < n; i++ {
		var aContainer Container
		aContainer.Capacity = cap
		arrContainer = append(arrContainer, aContainer)
	}

	for true {
		idx := getRandomNumber(int(n), int(1))
		arrContainer[idx-1].fill()
		if arrContainer[idx-1].isFull() {
			break
		}
	}

	fmt.Println(arrContainer)
}
