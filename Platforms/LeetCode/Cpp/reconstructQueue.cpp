class Solution {
public:
    vector<vector<int>> reconstructQueue(vector<vector<int>>& people) {
        int n = people.size();

        sort(people.begin(), people.end());

        vector<vector<int>> ans(n);

        for (int i = 0; i < n; i++) {
            int need = people[i][1];
            int height = people[i][0];

            for (int pos = 0; pos < n; pos++) {
                if (ans[pos].empty() && need==0) {
                        ans[pos] = people[i];
                        break;
                }
                else if (ans[pos].empty() || ans[pos][0] >= height ) {
                    need--;
                }
            }
        }

        return ans;
    }
};

