class Solution {
public:
    bool threeFibonacciSum(int n) {
        int a = 0, b = 1;
        
        while (true) {
            int sum = a + b;
            long long totalSum = sum + a + b;
            if (totalSum > n) return false;
            else if (totalSum == n) return true;
            a = b;
            b = sum;
        }

        return false;
    }
};