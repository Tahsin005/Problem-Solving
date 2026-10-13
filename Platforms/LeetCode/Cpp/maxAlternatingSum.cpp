class Solution {
public:
    long long maxAlternatingSum(vector<int>& nums) {
        int n = nums.size();
        vector<long long> sq(n);
        for (int i = 0; i < n; i++) {
            sq[i] = 1LL * nums[i] * nums[i];
        }
        sort(sq.rbegin(), sq.rend());
        long long ans = 0;
        int pos = (n + 1) / 2;
        for (int i = 0; i < n; i++) {
            if (i < pos) ans += sq[i];
            else ans -= sq[i];
        }
        return ans;
    }
};
