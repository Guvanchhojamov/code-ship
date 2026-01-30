package dp

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = int(1e9)
	}
	dp[0] = 0
	// [1,3,4,5]
	// [0,1,2,3,4,5,6,7]
	//  0,-,-,-,-,-,-,-
	for rem := 1; rem < amount+1; rem++ {
		for _, coin := range coins {
			if rem-coin >= 0 {
				dp[rem] = min(dp[rem], 1+dp[rem-coin])
			}
		}

	}
	if dp[amount] >= int(1e9) {
		return -1
	}
	return dp[amount]
}
