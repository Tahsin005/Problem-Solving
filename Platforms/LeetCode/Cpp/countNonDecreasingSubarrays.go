class Solution {
public:
    long long countNonDecreasingSubarrays(vector<int>& nums, int k) {
        int n = nums.size();
        long long ans = 0;
        long long operations = 0;
        priority_queue<pair<int, int>> q;
        vector<vector<int>> cnt(n);
        vector<int> pos(n, n);
        stack<int> st;

        for (int left = 0, right = 0; right < n; right++) {
            while (!st.empty() && nums[right] >= nums[st.top()]) {
                pos[st.top()] = right;
                st.pop();
            }
            if (!st.empty()) {
                cnt[st.top()].push_back(right);
            }
            st.push(right);

            q.emplace(nums[right], right);
            operations += max(0, q.top().first - nums[right]);

            while (operations > k) {
                for (int i : cnt[left]) {
                    operations -=
                        (long long)(nums[left] - nums[i]) *
                        (min(pos[i], right + 1) - i);
                }
                left++;
            }

            while (!q.empty() && q.top().second < left) {
                q.pop();
            }

            ans += right - left + 1;
        }

        return ans;
    }
};