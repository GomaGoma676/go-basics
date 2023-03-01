package main

import (
	"fmt"
	"time"
)

func main() {
	a := -1
	if a == 0 {
		fmt.Println("zero")
	} else if a > 0 {
		fmt.Println("positive")
	} else {
		fmt.Println("negative")
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// for {
	// 	fmt.Println("working")
	// 	time.Sleep(2 * time.Second)
	// }
	var i int
	for {
		if i > 3 {
			break
		}
		fmt.Println(i)
		i += 1
		time.Sleep(300 * time.Millisecond)
	}

loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 2:
			continue
		case 3:
			continue
		case 8:
			break loop
		default:
			fmt.Printf("%v ", i)
		}
	}
	fmt.Printf("\n")

	items := []item{
		{price: 10.},
		{price: 20.},
		{price: 30.},
	}
	for _, i := range items {
		i.price *= 1.1
	}
	fmt.Printf("%+v\n", items)
	for i := range items {
		items[i].price *= 1.1
	}
	fmt.Printf("%+v\n", items)

}

type item struct {
	price float32
}
