package main

func FirstNumber(number int) int {
	result := 0
	for i := number; i > 0; i /= 10 {
		result = i
	}
	return result
}
