class Solution {
public:
    string findLexSmallestString(string s, int a, int b) {
        int n = s.size();
        unordered_set<string> visited;
        queue<string> q;
        q.push(s);
        visited.insert(s);
        string ans = s;

        while (!q.empty()) {
            string cur = q.front();
            q.pop();

            if (cur < ans) ans = cur;

            string added = cur;
            for (int i = 1; i < n; i += 2) {
                added[i] = ((added[i] - '0' + a) % 10) + '0';
            }
            if (!visited.count(added)) {
                visited.insert(added);
                q.push(added);
            }

            string rotated = cur.substr(n - b) + cur.substr(0, n - b);
            if (!visited.count(rotated)) {
                visited.insert(rotated);
                q.push(rotated);
            }
        }

        return ans;
    }
};