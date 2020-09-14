package arraysandslices

func AddSum(arr []int) (sum int) {
	for _, num := range arr {
		sum += num
	}
	return
}

func ArrSums(arrays ...[]int) (arrOfSums []int) {

	for _, arr := range arrays {
		var arrSum int

		if len(arr) == 0 {
			arrSum = 0
		} else {
			arrSum = AddSum(arr)
		}
		arrOfSums = append(arrOfSums, arrSum)
	}

	return
}
