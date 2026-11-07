func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}
	//bucket sort
	bucket:=make([][]int,len(nums)+1)
	for n,f := range count{
		bucket[f]=append(bucket[f],n)
	}
	var res []int
	for f := len(bucket)-1; f>=0 && len(res)<k;f--{
		res = append(res,bucket[f]...)
	}
	return res
}
