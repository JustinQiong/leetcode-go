package leetcode

/*
1052. Grumpy Bookstore Owner
Sliding window solution:
Calculate the satisfied customers if
bookstore owner is not grumpy from 0 to minutes-1.
Then move the sliding window from i to the end, comparing the i-minutes
and i element of grumpy array. Recalculate the satisfied customers
and find the max satisfied customer number.
*/
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	res := 0
	for i := range minutes {
		res += customers[i]
	}

	n := len(customers)
	for i := minutes; i < n; i++ {
		if grumpy[i] == 0 {
			res += customers[i]
		}
	}

	mx := res
	for i := minutes; i < n; i++ {
		if grumpy[i-minutes] == 1 {
			res -= customers[i-minutes]
		}

		if grumpy[i] == 1 {
			res += customers[i]
		}

		mx = max(mx, res)
	}

	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
