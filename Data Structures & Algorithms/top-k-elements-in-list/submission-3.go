func topKFrequent(nums []int, k int) []int {
	hash:=make(map[int]int)
	for _,n := range nums{
		hash[n]++
	}
	key:=make([]int,0,len(hash))
	for k,_ := range hash{
		key=append(key,k)
	}
	sort.Slice(key,func(i,j int) bool {return hash[key[i]]>hash[key[j]]})
	return key[:k]
}
