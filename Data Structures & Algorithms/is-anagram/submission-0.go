func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sr := []rune(s)
	tr := []rune(t)
	sort.Slice(sr, func(i, j int) bool { return sr[i] < sr[j] })
	sort.Slice(tr, func(i, j int) bool { return tr[i] < tr[j] })
	for i := range sr {
		if sr[i] != tr[i] {
			return false
		}
	}
	return true
}
