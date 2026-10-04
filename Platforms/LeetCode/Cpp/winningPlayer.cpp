class Solution {
public:
    string winningPlayer(int x, int y) {
        int cnt = 0;
        while (x > 0 && y > 3) {
            x -= 1;
            y -= 4;
            cnt++;
        }
        return cnt % 2 ? "Alice" : "Bob";
    }
};
