/*
1997. First Day Where You Have Been in All the Rooms
dynamic programming solution:
The first day you in room i base on the first day you in i-1.
When we first arrive at room i-1, we have to revisit nextVisits[i-1].
And then move from room nextVisits[i-1] to room i-1.
So the transfer equation will be
f[i] = f[i-1]+1 + f[i-1] - f[nextVisit[i-1]] + 1
*/
func firstDayBeenInAllRooms(nextVisit []int) int {
    n := len(nextVisit)
    f := make([]int, n)
    const mod = 1e9 + 7
    for i := 1; i < n; i++ {
        f[i] = (f[i-1] + 1 + f[i-1] - f[nextVisit[i-1]] + 1 + mod) % mod
    }
    return f[n-1]
}