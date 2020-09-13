package integers

// Add takes two integers and returns the sum of them.
func Add(a, b int) int {
	return a + b
}

func AddSum(arr []int) (sum int) {
	for _, num := range arr {
		sum += num
	}
	return
}
