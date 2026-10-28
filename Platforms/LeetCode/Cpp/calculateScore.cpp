class Solution {
public:
    int calculateScore(int a, int b) {
        int u = __builtin_popcount(a);
        int v = __builtin_popcount(b);

        if (u < v)
            return v / u;
        return u * v;
    }
};