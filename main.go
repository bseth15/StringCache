package main

import (
	"fmt"

	"example.com/string-cache/collections"
)

func main() {
	cache := collections.NewStringCache()

	for {
		fmt.Print("Insert command: ")

		var input string
		fmt.Scan(&input)
		if input == "quit()" {
			return
		}

		res, err := cache.Call(input)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}
}
