package main

func findMin(nums []int) int {
	low, high := 0, len(nums) - 1

	//最小值一定落在 [low, high] 的区间中，如果当前已经满足 nums[low] < nums[high]，则可以直接返回 nums[low]
	if nums[low] < nums[high] {
		return nums[low]
	}

	//区间可以划分成两个排序好的子区间，每个子区间都是升序排序，而最小值是两个区间的第一个元素中较小的那一个
	for low < high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > nums[high] { //说明 mid 落在了左边递增区间内，最小值一定在 mid 的右边
			low = mid + 1
		} else {
			high = mid //注意这里不能写成 mid + 1，因为有可能此时 mid 刚好就是右边递增区间的第一个元素
		}
	}
	return nums[low]
}
