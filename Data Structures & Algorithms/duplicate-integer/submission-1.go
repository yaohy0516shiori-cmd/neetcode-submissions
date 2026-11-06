func hasDuplicate(nums []int) bool {
    hash:=make(map[int]bool)
	for _,n:=range nums{
		if hash[n]{
			return true
		}
		hash[n]=true
	}
	return false
}
