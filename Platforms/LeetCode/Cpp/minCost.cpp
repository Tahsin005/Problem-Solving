class Solution {
public:
    int minCost(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        vector<vector<bitset<1024>>> dp(m, vector<bitset<1024>>(n));

        dp[m - 1][n - 1][grid[m - 1][n - 1]] = 1;

        for (int i = m - 1; i >= 0; i--)
            for (int j = n - 1; j >= 0; j--) {
                if (i == m - 1 && j == n - 1) continue;

                if (i + 1 < m)
                    for (int x = 0; x < 1024; x++)
                        if (dp[i + 1][j][x]) dp[i][j][x ^ grid[i][j]] = 1;

                if (j + 1 < n)
                    for (int x = 0; x < 1024; x++)
                        if (dp[i][j + 1][x]) dp[i][j][x ^ grid[i][j]] = 1;
            }

        for (int x = 0; x < 1024; x++)
            if (dp[0][0][x]) return x;

        return 0;
    }
};