func productExceptSelf(nums []int) []int {
	n := len(nums)
	// O(1) space extra
	res := make([]int, n)
	p := 1
	for i := 0; i < n; i++ {
		res[i] = p
		p *= nums[i]
	}
	p = 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= p
		p *= nums[i]
	}
	return res
}
