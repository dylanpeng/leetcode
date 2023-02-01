package t2546

// 同时存在1或都为0
func makeStringsEqual(s string, target string) bool {
	s1, t1 := false, false
	for i := 0; i < len(s); i++ {
		if s[i] == 49 && !s1 {
			s1 = true
		}

		if target[i] == 49 && !t1 {
			t1 = true
		}

		if s1 && t1 {
			return true
		}
	}

	if !s1 && !t1 {
		return true
	}

	return false
}
