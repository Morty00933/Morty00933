package main

import "fmt"

func main() {
	var i = 500000001
	var count = 0
	for count < 5 {
		var halfI = i / 2
		var del = 1
		var countDel = 0
		var j = 2
		for j <= halfI {
			if i%j == 0 {
				countDel++
				del = del * j
				if del > i {
					break
				} else if countDel == 5 {
					fmt.Println(del)
					count++
					break
				}
			}
			j++
		}
		i++
	}
}
