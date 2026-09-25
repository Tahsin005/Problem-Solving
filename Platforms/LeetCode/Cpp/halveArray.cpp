class Solution {
public:
    int halveArray(vector<int>& nums) {
        priority_queue<double> pq(nums.begin(), nums.end());
        double sum = 0;
        int c  =0;
        for (int i: nums) sum += i;

        double a = sum / 2;
        while (sum > a) {
            double t = pq.top();
            pq.pop();
            double s = t/2;
            pq.push(s);
            sum = sum + s - t;
            c++;
        }
        
        return c;
    }
};