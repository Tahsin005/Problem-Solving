class Solution {
public:
    vector<int> rearrangeArray(vector<int>& nums) {
        map<int, int> freq;
        for (int x : nums) freq[x]++;
        vector<int> ans;
        while (!freq.empty()) {
            for (auto& [x, f] : freq) {
                ans.push_back(x);
                --f;
            }
            erase_if(freq, [](const auto& p) { return p.second == 0; });
        }
        return ans;
    }
};