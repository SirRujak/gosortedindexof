package gosortedindexof

func indexOfString(left []string, right []string) []int {
	var result []int
	result = make([]int, len(right), len(right))
	i := 0
	j := 0

	var a, b string
	for i < len(left) && j < len(right) {
		a = left[i]
		b = right[j]

		if a == b {
			result[j] = i
			j++
			continue
		}

		if a < b {
			i++
			continue
		}

		result[j] = -1
		continue
	}

	for ; j < len(right); j++ {
		result[j] = -1
	}

	return result
}

func indexOfInt(left []int, right []int) []int {
	var result []int
	result = make([]int, len(right), len(right))
	i := 0
	j := 0

	var a, b int
	for i < len(left) && j < len(right) {
		a = left[i]
		b = right[j]

		if a == b {
			result[j] = i
			j++
			continue
		}

		if a < b {
			i++
			continue
		}

		result[j] = -1
		continue
	}

	for ; j < len(right); j++ {
		result[j] = -1
	}

	return result
}
