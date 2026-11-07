func groupAnagrams(strs []string) [][]string {
	sortkey:= func (s string) string{
		b:=[]byte(s)
		sort.Slice(b,func(i,j int) bool {return b[i]<b[j]})
		return string(b)
	}
	group := make(map[string][]string)
	for _,s := range strs{
		k:=sortkey(s)
		group[k]=append(group[k],s)
	}
	res:=make([][]string,0,len(group))
	for _,v := range group{
		res=append(res,v)
	}
	return res
}
