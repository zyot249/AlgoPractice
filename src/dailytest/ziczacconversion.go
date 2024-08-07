package dailytest

func convert(s string, numRows int) string {
	result := make([]byte, len(s))
	idx := 0
	for row := 0; row < numRows; row++ {
		count := 0
		cur, next := -1, row
		for next < len(s) {
			if next != cur {
				result[idx] = s[next]
				idx++
				cur = next
			}

			if count%2 == 0 {
				next += (numRows - row - 1) * 2
			} else {
				next += row * 2
			}

			count++
		}
	}

	return string(result)
}
