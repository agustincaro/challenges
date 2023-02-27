package main

import "fmt"

func main() {
	fmt.Println("Ej 1")
}

func listAdds(numbers []int, k int) bool {

	length := len(numbers)
	if length == 2 {
		return numbers[0]+numbers[1] == k
	}

	for i := 0; i < length; i++ {
		sublist := numbers[1:]
		if subListAdds(sublist, numbers[i], k) {
			return true
		}
	}

	return false
}

func subListAdds(sublist []int, number int, k int) bool {

	for i := 0; i < len(sublist); i++ {
		if sublist[i]+number == k {
			return true
		}
	}
	return false
}
