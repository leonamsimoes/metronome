package main

import (
	"fmt"
	"time"
)

func main() {
	start(100, 4)
}

func start(bpm float64, comp int) {
	d := time.Duration(float64(time.Minute) / bpm)
	fmt.Printf("Time:%d \n", d)
	var i = 1
	for {
		time.Sleep(d)
		if i < comp {
			fmt.Printf(" [bip %d]", i)
		}
		if i == comp {
			fmt.Printf(" [BIP %d]\n", i)
			i = 0
		}
		i++
	}
}
