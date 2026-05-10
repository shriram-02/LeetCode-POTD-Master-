class Solution {
    public int maximumJumps(int[] nums, int target) {
        int n = nums.length;

        int[] dp = new int[n];

        // -1 means unreachable
        for (int i = 0; i < n; i++) {
            dp[i] = -1;
        }

        dp[0] = 0;

        for (int i = 1; i < n; i++) {

            for (int j = 0; j < i; j++) {

                long diff = (long) nums[i] - nums[j];

                if (diff >= -target && diff <= target && dp[j] != -1) {
                    dp[i] = Math.max(dp[i], dp[j] + 1);
                }
            }
        }

        return dp[n - 1];
    }
}