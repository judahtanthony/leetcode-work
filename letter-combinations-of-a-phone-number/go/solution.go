package leetcode

func letters(n byte) string {
	switch n {
	case '2':
		return "abc"
	case '3':
		return "def"
	case '4':
		return "ghi"
	case '5':
		return "jkl"
	case '6':
		return "mno"
	case '7':
		return "pqrs"
	case '8':
		return "tuv"
	case '9':
		return "wxyz"
	default:
		return ""
	}
}

func letterCombinationsRealloc(digits string) []string {
	n := len(digits)
	if n == 0 {
		return []string{}
	}

	results := make([]string, 0, 3)
	for _, c := range letters(digits[0]) {
		if n == 1 {
			results = append(results, string(c))
		} else {
			for _, str := range letterCombinations(digits[1:]) {
				results = append(results, string(c)+str)
			}
		}
	}

	return results
}

func lc(buff []string, add byte) []string {
	n, res := len(buff), buff[:]

	ns := letters(add)
	for i, nsn := 0, len(ns); i < nsn; i++ {
		s := string(ns[i])
		if n == 0 {
			res = append(res, s)
		} else {
			for i2, s2 := range buff {
				if i == nsn-1 {
					res[i2] = s2 + s
				} else {
					res = append(res, s2+s)
				}
			}
		}
	}

	return res
}

// Reduce the number of reallocations
func letterCombinationsReduced(digits string) []string {
	n := len(digits)
	if n == 0 {
		return []string{}
	}
	cap, exp := 3, n-1
	for exp > 0 {
		cap *= 3
		exp--
	}

	results := make([]string, 0, cap)
	for i := 0; i < n; i++ {
		results = lc(results, digits[i])
	}

	return results
}

func calcKey(digits string) ([][]byte, int) {
	n := len(digits)
	ret := make([][]byte, 0, n)
	m := 1

	for i := 0; i < n; i++ {
		chars := letters(digits[i])
		nn := len(chars)
		if nn > 0 {
			m *= nn
			ret = append(ret, []byte(chars))
		}
	}

	if len(ret) == 0 {
		m = 0
	}

	return ret, m
}

// Fully tabularized
func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return []string{}
	}
	key, m := calcKey(digits)
	n = len(key)

	// Strings are immutable, so we need to build our
	// table as a grid of bytes.
	tab := make([][]byte, m)
	stepm := m
	// For each digit...
	for col := 0; col < n; col++ {
		chars := key[col]
		charsn := len(key[col])
		stepm /= charsn
		// ...assign a character per permutation.
		for row := 0; row < m; row++ {
			if col == 0 {
				tab[row] = make([]byte, n)
			}
			tab[row][col] = chars[(row/stepm)%charsn]
		}
	}
	// Now we need to turn that grid back into a list of strings.
	ret := make([]string, m)
	for i, bytes := range tab {
		ret[i] = string(bytes)
	}

	return ret
}
