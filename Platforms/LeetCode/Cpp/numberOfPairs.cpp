class Solution {
public:
    vector<int> numberOfPairs(vector<int>& nums) {
        int freq[101] = {0};
        for (int x : nums) {
            freq[x]++;
        }

        int pairs = 0;
        int leftovers = 0;

        for (int i = 0; i < 101; i++) {
            if (freq[i] > 0) {
                pairs += freq[i] / 2;   
                leftovers += freq[i] % 2;
            }
        }

        return {pairs, leftovers};
    }
};