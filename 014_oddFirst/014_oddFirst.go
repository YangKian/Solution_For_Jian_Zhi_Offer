package main

//要点：考虑泛化能力，设置一个函数类型func(int) bool的变量用来定义排序的规则，
//这样一来无论是求奇数在偶数前还是划分能被3整除 之类的数都可以使用同一函数执行，
//只需要设置特定的排序函数即可
type transfer func(int) bool

func Reorder(nums *[]int, flag transfer) {
	length := len(*nums)

	if length == 0 {
		return
	}

	left, right := 0, length - 1
	for left < right {
		for left < right && !flag((*nums)[left]) {
			left++
		}

		for left < right && flag((*nums)[right]) {
			right--
		}

		if left < right {
			(*nums)[left], (*nums)[right] = (*nums)[right], (*nums)[left]
		}
	}
}

func isEven(num int) bool {
	return (num & 1) == 0
}

//func main() {
//	nums := []int{0, 1, 2, 3, 4, 5}
//	Reorder(&nums,isEven)
//	fmt.Println(nums)
//}