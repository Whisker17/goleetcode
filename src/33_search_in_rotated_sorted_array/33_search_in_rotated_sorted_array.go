package main

/*func search(nums []int, target int) int {
	left,right := 0, len(nums)-1
	for left <= right {
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		left++
		right--
	}
	return -1
}*/

func search(nums []int, target int) int {
	left,right := 0, len(nums)-1
	for left <= right {
		mid :=  (left+right)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && nums[mid] > target {
				right = mid-1
			} else {
				left = mid+1
			}
		}
		if nums[right] >= nums[mid] {
			if target > nums[mid] && target <= nums[right] {
				left = mid+1
			} else {
				right = mid-1
			}
		}
	}
	return -1
}
