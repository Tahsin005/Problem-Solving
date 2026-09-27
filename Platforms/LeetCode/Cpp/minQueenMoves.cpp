class Solution {
public:
    int minQueenMoves(vector<int>& s, vector<int>& t) {
        if (s[0] == t[0] && s[1] == t[1]) return 0;
        if (s[0] == t[0] || s[1] == t[1] || abs(s[0] - t[0]) == abs(s[1] - t[1])) return 1;
        return 2;
    }
};
