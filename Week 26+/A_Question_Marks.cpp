#include<bits/stdc++.h>
using namespace std;

int main () {
    int _; cin >> _; while (_--) {
        int n;
        cin >> n;
        string s; cin >> s;

        map<char, int> mp;
        for (char ch : s) {
            mp[ch]++;
        }

        int ans = 0;
        
        for (auto [x, y] : mp) {
            if (x != '?') {
                ans += min(y, n);
            }
        }

        cout << ans << '\n';
    }
    return 0;
}