#include <bits/stdc++.h>
using namespace std;
#define ll long long
const int M = 1e18+7;

void solve() {
    string s;
    cin >> s;
    
    sort(s.begin(), s.end());
    bool ok = true;
    for (int i = 1; i < s.length(); i++) {
        if (s[i] - s[i - 1] != 1) {
            ok &= false;
        }
    }
    
    cout << (ok ? "YES": "NO") << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt = 1; cin >> tt; while (tt--) solve();
}
