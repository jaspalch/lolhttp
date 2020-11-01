package server

// find checks whether string s exists in
// slice strs
func find(s string, strs []string) bool {
	found := false
	for _, str := range strs {
		if str == s {
			found = true
		}
	}
	return found
}
