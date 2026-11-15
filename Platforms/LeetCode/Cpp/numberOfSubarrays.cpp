class Solution {
public:
    int numOfSubarrays(vector<int>& arr) {
        const int MOD = 1e9 + 7;
        int n = arr.size();
        
        vector<vector<int>> dp(2, vector<int>(2, 0));
        
        int count = dp[1][0]; 
        for (int i = 0; i < n; i++) {
            unsigned char idx = i & 1; 
            unsigned char parity = arr[i] & 1; 
            
            dp[parity][idx] = (1 + dp[0][!idx]) % MOD;
            dp[!parity][idx] = dp[1][!idx] % MOD;
            
            count = (count + dp[1][idx]) % MOD;
        }
        return count;
    }
};
