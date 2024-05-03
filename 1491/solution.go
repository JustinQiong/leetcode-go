package leetcode

func average(salary []int) float64 {
	n := len(salary)
	mx := salary[0]
	mn := salary[0]
	total := 0
	for _, x := range salary {
		mx = max(x, mx)
		mn = min(x, mn)
		total += x
	}
	c := float64(n - 2)
	t := float64(total - mx - mn)
	return t / c
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
