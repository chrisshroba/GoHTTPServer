package main

import "time"

func main() {
	defer func() {
		println("DEFERRED MUTHAFUCKA")
	}()
	i := 0
	for {
		i++
		println("Hi", i)

		time.Sleep(100 * time.Millisecond)
	}
}
