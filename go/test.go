package main

import (
    "fmt"
    "time"
)


func main()  {
	fmt.Println("hello world")

	// LOOP : FOR
	i:=0
	for i<=3{
		fmt.Println(i)
		i=i+1
	}

    for i := time.Sunday; i <= time.Saturday; i++ {
        fmt.Println(i)
    }
	
}

// go build test.go :
    //Sometimes we’ll want to build our programs into binaries.
	//We can then execute the built binary directly.

