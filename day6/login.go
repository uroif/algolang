package main

import "fmt"

func main() {
	var n int
	var user, pass string
	db := make(map[string]string)

	fmt.Scan(&n)
	for n > 0 {
		n--
		fmt.Scan(&user, &pass)
		db[user] = pass
	}

	fmt.Scan(&n)
	for n > 0 {
		n--
		fmt.Scan(&user, &pass)
		if pass == db[user] {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}
	}
}
