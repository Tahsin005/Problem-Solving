class Solution {
public:
    long long maximizeSum(vector<int>& nums, int k) {
        long long n = nums.size(), sum = 0;
        priority_queue<long long> pq;
        for (int i = 0; i < n; i++) {
            string s = to_string(nums[i]);
            sum += nums[i];
            sort(s.rbegin(), s.rend());    

            long long new_num = stoi(s);
            long long diff = new_num - nums[i];
            pq.push(diff);
        }

        for (int i = 0; i < k; i++) {
            sum += pq.top();
            pq.pop();
        }

        return sum;   
    }
};
