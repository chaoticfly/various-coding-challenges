package main

//TwoSum is not efficient
// func TwoSum(arrNums []int, total int) [2]int {
// 	var i1, i2 int
// 	for i := 0; i < len(arrNums)-1; i++ {
// 		for j := i + 1; j < len(arrNums); j++ {
// 			tot := arrNums[i] + arrNums[j]
// 			if tot == total {
// 				i1 = i
// 				i2 = j
// 				break
// 			}
// 		}
// 	}
// 	return [2]int{i1, i2}
// }
//TwoSum More Efficient
func TwoSum(arrNums []int, total int) [2]int {
	j := 1
	i := 0
	for i = 0; i < len(arrNums)-1; i++ {
		num := total - arrNums[i]
		status, ind := find(arrNums, num)
		if status {
			j = ind
			break
		}
	}
	return [2]int{i, j}
}

func find(slice []int, elem int) (bool, int) {
	for i := len(slice) - 1; i > -1; i-- {
		if slice[i] == elem {
			return true, i
		}
	}
	return false, -1
}

// func main() {
// 	//arr := []int{2, 4, 5, 7, 9, 1}
// 	fmt.Println(TwoSum([]int{2, 2, 3}, 4))

// }

//input is [2, 4, 5, 7,9, 1] - 14
