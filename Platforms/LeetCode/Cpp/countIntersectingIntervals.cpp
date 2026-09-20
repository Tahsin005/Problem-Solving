class Solution {
public:
    int countIntersectingIntervals(vector<vector<int>>& intervals) {
        int n = intervals.size();
        int count = 0;
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                if (intervals[i][0] <= intervals[j][1] &&
                    intervals[j][0] <= intervals[i][1]) {
                    count++;
                }
            }
        }
        return count;
    }
};