package leetcode

/*
2739. Total Distance Traveled
The fuel can be filled by additionalTank will be
the minimum value between addtionalTank and (mainTank-5)/4 + 1.
So the distance traveled will be
10 * (mainTank + min((mainTank-1)/4, additionalTank))
*/
func distanceTraveled(mainTank int, additionalTank int) int {
	return 10 * (mainTank + min((mainTank-1)/4, additionalTank))
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
