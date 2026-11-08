class Solution {
public:
    long long countValidSubstrings(string s) {
        long long res = 0, len = 0, n = s.size();
        for (int i = 0; i < n; ++i) {
            ++len;
            if (i >= 2 && s.substr(i - 2, 3) == "abc") {
                res += (len - 1) * len / 2 - 1;
                len = 2;
            }
        }
        
        return res + len * (len + 1) / 2;
    }
};
