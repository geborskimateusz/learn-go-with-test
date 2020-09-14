package arraysandslices

import (
	"errors"
)

func AddSum(arr []int) (sum int) {
	for _, num := range arr {
		sum += num
	}
	return
}

func ArrSums(arrays ...[]int) (arrOfSums []int, error) {

	for _, arr := range arrays {
		if(len(arr) == 0) {
			return errors.New("Array must be initialized")
		}
		arrSum := AddSum(arr)
		arrOfSums = append(arrOfSums, arrSum)
	}

	return
}
