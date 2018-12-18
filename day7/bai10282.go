package main

import "fmt"

func main() {
	var n int
	var eng, chingchong string
	dict := make(map[string]string)

	fmt.Scan(&n)
	for n > 0 {
		n--
		fmt.Scan(&eng, &chingchong)
		dict[chingchong] = eng
	}

	fmt.Println(dict)

	fmt.Scan(&n)
	for n > 0 {
		n--
		fmt.Scan(&eng)

		str, boo := isExisted(dict, eng)
		if boo == true {
			fmt.Println(str)
		} else {
			fmt.Println("eh")
		}
	}
}

func isExisted(m map[string]string, s string) (string, bool) {
	for k, v := range m {
		if k == s {
			return v, true
		}
	}
	return "", false
}
