package main

import "fmt"

func main() {

	users := map[string]int{
		"user1": 1,
		"user2": 2,
	}
	for k, v := range users {
		fmt.Println(k, v)
	}
}
