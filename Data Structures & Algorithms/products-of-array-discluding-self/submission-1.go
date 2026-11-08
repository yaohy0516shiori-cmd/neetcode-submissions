func productExceptSelf(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	suf := make([]int, n)
	p := 1
	for i := 0; i < n; i++ {
		pre[i] = p
		p *= nums[i]
	}
	p = 1
	for i := n - 1; i >= 0; i-- {
		suf[i] = p
		p *= nums[i]
	}
	res := make([]int, n)
	for i := range nums {
		res[i] = pre[i] * suf[i]
	}
	return res
}