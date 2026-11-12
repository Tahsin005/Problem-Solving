class Solution {
public:
    bool canRemoveAll(vector<int>& types, vector<int>& colors, vector<int>& groupSize) {
        vector<array<int, 3>> s;
        for (int i = 0; i < types.size(); i++) {
            int t = types[i], c = colors[i];
            if (!s.empty() && s.back()[0] == t && s.back()[1] == c) {
                s.back()[2]++;
            } else {
                s.push_back({t, c, 1});
            }
            if (s.back()[2] == groupSize[t]) {
                s.pop_back();
            }
        }
        return s.empty();
    }
};
