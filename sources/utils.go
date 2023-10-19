package sources

func bInt64(b []byte) int64 {
	var res int64
	var factor int64 = 1
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] >= 48 && b[i] <= 57 {
			res += int64(b[i]-48) * factor
			factor *= 10
		}
	}
	return res
}

func readInts(line string, n int) []int64 {
	const intLen = 15
	var values []int64
	buf := []byte(line)
	currentNumber := make([]byte, 0, intLen)
	for i := 0; i < len(buf); i++ {
		if buf[i] >= 48 && buf[i] <= 57 {
			currentNumber = append(currentNumber, buf[i])
		} else if len(currentNumber) > 0 {
			values = append(values, bInt64(currentNumber))
			if len(values) == n {
				return values
			}
			currentNumber = make([]byte, 0, intLen)
		}
	}
	if len(currentNumber) > 0 {
		values = append(values, bInt64(currentNumber))
	}
	return values
}
