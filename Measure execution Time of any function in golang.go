package main

import (
	"time"
	"log"
)

func main() {
	loop(1, 5) 
}

func loop(x, y int) {
	// defer statement executes just before the return 
	defer funcTime(time.Now(), "loop")
	for i := x; i <= y; i++ {
		log.Println(i)
	}
}

/* Invoke funcTime() in the beginning of any function 
   to measure function time using defer statement */

func funcTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s function took %s", name, elapsed)
}
