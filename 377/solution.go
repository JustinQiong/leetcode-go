package leetcode

/*
377. Combination Sum IV
Dynamic Programming Solution:
Define a dynamic programming array dp[].
Each element in dp stands for the combination count
of sum value x. Iterate through each element of num in nums and accumulate
the combination count base on dp[i-num].
Return the dp[target].
*/
func combinationSum4(nums []int, target int) int {
	f := make([]int, target+1)
	f[0] = 1
	for i := 1; i <= target; i++ {
		for _, x := range nums {
			if x <= i {
				f[i] += f[i-x]
			}
		}
	}

	return f[target]
}
